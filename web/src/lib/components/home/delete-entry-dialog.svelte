<script lang="ts">
	import { Button, buttonVariants } from "$lib/components/ui/button";
	import * as Dialog from "$lib/components/ui/dialog";
	import { createMutation, useQueryClient } from "@tanstack/svelte-query";
	import { API_URL } from "$lib/queries/api";
	import { toast } from "svelte-sonner";
	import { Trash2, Loader2 } from "@lucide/svelte";

	const { key }: { key: string } = $props();
	let open = $state(false);
	const queryClient = useQueryClient();

	const createDeleteMutation = () =>
		createMutation(() => ({
			mutationFn: async (keyStr: string) => {
				const res = await fetch(`${API_URL}/del/${keyStr}`, { method: "DELETE" });
				if (!res.ok) throw new Error("deletion failed");
				return res.json();
			},
			onSuccess: () => {
				queryClient.invalidateQueries({ queryKey: ["getAllData"] });
				queryClient.invalidateQueries({ queryKey: ["stats"] });
				toast.success("entry deleted");
				open = false;
			},
			onError: () => {
				toast.error("couldn't delete");
			},
		}));

	const deleteMutation = createDeleteMutation();
</script>

<Dialog.Root bind:open>
	<Dialog.Trigger
		class={buttonVariants({ variant: "ghost", size: "icon-sm" }) +
			' text-muted-foreground hover:text-destructive'}
		title="delete"
		aria-label="delete"
	>
		<Trash2 class="size-3.5" />
	</Dialog.Trigger>
	<Dialog.Content class="sm:max-w-[380px]">
		<Dialog.Header>
			<Dialog.Title class="font-serif text-xl">delete entry?</Dialog.Title>
			<p class="text-xs text-muted-foreground">
				<span class="font-mono">{key}</span> will be removed from the store.
			</p>
		</Dialog.Header>
		<Dialog.Footer class="gap-2">
			<Dialog.Close class={buttonVariants({ variant: "ghost", size: "sm" })}>
				cancel
			</Dialog.Close>
			<Button
				variant="destructive"
				size="sm"
				disabled={deleteMutation.isPending}
				onclick={() => deleteMutation.mutateAsync(key)}
			>
				{#if deleteMutation.isPending}
					<Loader2 class="size-3.5 animate-spin" />
				{:else}
					<Trash2 class="size-3.5" />
				{/if}
				delete
			</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
