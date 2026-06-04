<script lang="ts">
	import { Button } from "$lib/components/ui/button";
	import { createMutation, useQueryClient } from "@tanstack/svelte-query";
	import { API_URL } from "$lib/queries/api";
	import { toast } from "svelte-sonner";
	import { Loader2, Sprout } from "@lucide/svelte";

	const queryClient = useQueryClient();

	const seedMutation = createMutation(() => ({
		mutationFn: async () => {
			const res = await fetch(`${API_URL}/seed`, { method: "POST" });
			if (!res.ok) throw new Error("seed failed");
			return res.json();
		},
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ["getAllData"] });
			queryClient.invalidateQueries({ queryKey: ["stats"] });
			toast.success("seeded sample data");
		},
		onError: () => {
			toast.error("couldn't seed");
		},
	}));
</script>

<Button
	variant="ghost"
	size="sm"
	class="text-muted-foreground"
	disabled={seedMutation.isPending}
	onclick={() => seedMutation.mutate()}
>
	{#if seedMutation.isPending}
		<Loader2 class="size-3.5 animate-spin" />
	{:else}
		<Sprout class="size-3.5" />
	{/if}
	seed
</Button>
