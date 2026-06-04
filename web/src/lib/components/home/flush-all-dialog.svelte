<script lang="ts">
	import { Button, buttonVariants } from "$lib/components/ui/button";
	import * as Dialog from "$lib/components/ui/dialog";
	import { createMutation, useQueryClient } from "@tanstack/svelte-query";
	import { API_URL } from "$lib/queries/api";
	import { toast } from "svelte-sonner";
	import { Loader2, FlameKindling } from "@lucide/svelte";

	let open = $state(false);
	const queryClient = useQueryClient();

	const createFlushMutation = () =>
		createMutation(() => ({
			mutationFn: async () => {
				const res = await fetch(`${API_URL}/flushall`, { method: "DELETE" });
				if (!res.ok) throw new Error("flush failed");
				return res.json();
			},
			onSuccess: () => {
				queryClient.invalidateQueries({ queryKey: ["getAllData"] });
				queryClient.invalidateQueries({ queryKey: ["stats"] });
				toast.success("store flushed");
				open = false;
			},
			onError: () => {
				toast.error("couldn't flush");
			},
		}));

	const flushMutation = createFlushMutation();
</script>

<Dialog.Root bind:open>
	<Dialog.Trigger class={buttonVariants({ variant: "ghost", size: "sm" }) + ' text-muted-foreground hover:text-destructive'}>
		<FlameKindling class="size-3.5" />
		flush
	</Dialog.Trigger>
	<Dialog.Content class="sm:max-w-[380px]">
		<Dialog.Header>
			<Dialog.Title class="font-serif text-xl">flush everything?</Dialog.Title>
			<p class="text-xs text-muted-foreground">
				wipes every key from the store. cannot be undone.
			</p>
		</Dialog.Header>
		<Dialog.Footer class="gap-2">
			<Dialog.Close class={buttonVariants({ variant: "ghost", size: "sm" })}>
				cancel
			</Dialog.Close>
			<Button
				variant="destructive"
				size="sm"
				disabled={flushMutation.isPending}
				onclick={() => flushMutation.mutateAsync()}
			>
				{#if flushMutation.isPending}
					<Loader2 class="size-3.5 animate-spin" />
				{/if}
				flush all
			</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
