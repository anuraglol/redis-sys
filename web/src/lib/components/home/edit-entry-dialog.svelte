<script lang="ts">
	import { Button, buttonVariants } from "$lib/components/ui/button";
	import * as Dialog from "$lib/components/ui/dialog";
	import { Input } from "$lib/components/ui/input";
	import { Label } from "$lib/components/ui/label";
	import { API_URL } from "$lib/queries/api";
	import { createMutation, useQueryClient } from "@tanstack/svelte-query";
	import { Pencil } from "@lucide/svelte";
	import { toast } from "svelte-sonner";
	import { untrack } from "svelte";

	const { key, value, expiry }: { key: string; value: string; expiry: number } = $props();

	const queryClient = useQueryClient();

	let open = $state(false);
	let formValue = $state(untrack(() => value));
	let formTtl = $state(untrack(() => (expiry === -1 ? "" : expiry.toString())));

	const createSetMutation = () =>
		createMutation(() => ({
			mutationFn: async (payload: { key: string; value: string; ttl: string }) => {
				const body: { key: string; value: string; ex?: number } = {
					key: payload.key,
					value: payload.value,
				};
				if (payload.ttl) body.ex = parseInt(payload.ttl, 10);
				const res = await fetch(`${API_URL}/set`, {
					method: "POST",
					headers: { "Content-Type": "application/json" },
					body: JSON.stringify(body),
				});
				if (!res.ok) throw new Error("failed to save entry");
				return res.json();
			},
			onSuccess: () => {
				queryClient.invalidateQueries({ queryKey: ["getAllData"] });
				queryClient.invalidateQueries({ queryKey: ["stats"] });
				toast.success("entry updated");
				open = false;
			},
			onError: (e) => {
				toast.error(e.message);
			},
		}));

	const setMutation = createSetMutation();

	async function handleSubmit(e: SubmitEvent) {
		e.preventDefault();
		await setMutation.mutateAsync({ key, value: formValue, ttl: formTtl });
	}
</script>

<Dialog.Root bind:open>
	<Dialog.Trigger
		class={buttonVariants({ variant: "ghost", size: "icon-sm" })}
		title="edit"
		aria-label="edit"
	>
		<Pencil class="size-3.5" />
	</Dialog.Trigger>
	<Dialog.Content class="sm:max-w-[400px]">
		<Dialog.Header>
			<Dialog.Title class="font-serif text-xl">edit entry</Dialog.Title>
			<p class="text-xs text-muted-foreground font-mono">{key}</p>
		</Dialog.Header>

		<form onsubmit={handleSubmit} class="grid gap-4">
			<div class="grid gap-1.5">
				<Label for="form-value" class="text-xs text-muted-foreground">value</Label>
				<Input id="form-value" required bind:value={formValue} />
			</div>

			<div class="grid gap-1.5">
				<Label for="form-ttl" class="text-xs text-muted-foreground">
					ttl <span class="text-muted-foreground/60">(seconds)</span>
				</Label>
				<Input id="form-ttl" type="number" bind:value={formTtl} placeholder="leave empty for ∞" />
			</div>

			<Dialog.Footer class="mt-1 gap-2">
				<Dialog.Close class={buttonVariants({ variant: "ghost", size: "sm" })}>
					cancel
				</Dialog.Close>
				<Button type="submit" size="sm" disabled={setMutation.isPending}>
					{setMutation.isPending ? "saving…" : "save"}
				</Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
