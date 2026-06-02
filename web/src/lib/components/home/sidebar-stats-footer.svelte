<script lang="ts">
	import { Check, Flame } from "@lucide/svelte";
	import type { StatsQuery } from "$lib/queries/queries";

	let { statsQuery }: { statsQuery: StatsQuery } = $props();
</script>

<div
	class="p-4 border-t border-border bg-muted/20 font-mono text-[11px] space-y-1.5 text-muted-foreground"
>
	<div class="flex justify-between items-center">
		<span class="flex items-center gap-1.5"
			><Flame class="w-3.5 h-3.5 text-orange-500" /> Total Triggers:</span
		>
		<span class="text-foreground font-semibold"
			>{statsQuery.data?.total_commands || 0}</span
		>
	</div>
	<div class="flex justify-between items-center">
		<span class="flex items-center gap-1.5"
			><Check class="w-3.5 h-3.5 text-emerald-500" /> Operational Matrix:</span
		>
		<span class="text-emerald-500 font-bold uppercase tracking-tight text-[10px]"
			>Active</span
		>
	</div>
</div>
