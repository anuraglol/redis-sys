<script lang="ts">
	import { Button, buttonVariants } from "$lib/components/ui/button";
	import * as Dialog from "$lib/components/ui/dialog";
	import { Input } from "$lib/components/ui/input";
	import { Label } from "$lib/components/ui/label";
	import { API_URL } from "$lib/queries/api";
	import { createMutation, useQueryClient } from "@tanstack/svelte-query";
	import { Plus } from "@lucide/svelte";
	import { toast } from "svelte-sonner";

	const queryClient = useQueryClient();

	let open = $state(false);
	let key = $state("");
	let formValue = $state("");
	let formTtl = $state("");

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
				if (!res.ok) throw new Error("failed to write entry");
				return res.json();
			},
			onSuccess: () => {
				queryClient.invalidateQueries({ queryKey: ["getAllData"] });
				queryClient.invalidateQueries({ queryKey: ["stats"] });
				toast.success("entry saved");
				key = "";
				formValue = "";
				formTtl = "";
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
	<Dialog.Trigger class={buttonVariants({ variant: "default", size: "sm" })}>
		<Plus class="size-3.5" />
		add entry
	</Dialog.Trigger>
	<Dialog.Content class="sm:max-w-[400px]">
		<Dialog.Header>
			<Dialog.Title class="font-serif text-xl">add entry</Dialog.Title>
		</Dialog.Header>

		<form onsubmit={handleSubmit} class="grid gap-4">
			<div class="grid gap-1.5">
				<Label for="form-key" class="text-xs text-muted-foreground">key</Label>
				<Input id="form-key" required bind:value={key} placeholder="user:1" />
			</div>

			<div class="grid gap-1.5">
				<Label for="form-value" class="text-xs text-muted-foreground">value</Label>
				<Input id="form-value" required bind:value={formValue} placeholder="hello" />
			</div>

			<div class="grid gap-1.5">
				<Label for="form-ttl" class="text-xs text-muted-foreground">
					ttl <span class="text-muted-foreground/60">(seconds, optional)</span>
				</Label>
				<Input id="form-ttl" type="number" bind:value={formTtl} placeholder="3600" />
			</div>

			<Dialog.Footer class="mt-1 gap-2">
				<Dialog.Close class={buttonVariants({ variant: "ghost", size: "sm" })}>
					cancel
				</Dialog.Close>
				<Button type="submit" size="sm" disabled={setMutation.isPending}>
					{setMutation.isPending ? "saving…" : "save entry"}
				</Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
