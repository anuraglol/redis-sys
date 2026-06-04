<script lang="ts">
	import type { StatsQuery } from "$lib/queries/queries";

	let { statsQuery }: { statsQuery: StatsQuery } = $props();

	const rows = $derived([
		{ label: "commands", value: statsQuery.data?.total_commands ?? "0" },
		{ label: "get", value: statsQuery.data?.get_cmds ?? "0" },
		{ label: "set", value: statsQuery.data?.set_cmds ?? "0" },
		{ label: "expire", value: statsQuery.data?.expire_cmds ?? "0" },
	]);
</script>

<div class="m-3 mt-2 rounded-2xl border border-sidebar-border bg-card/60 backdrop-blur-sm p-4">
	<div class="flex items-center justify-between mb-3">
		<span class="text-[10px] font-medium uppercase tracking-[0.12em] text-muted-foreground">
			stats
		</span>
		<span class="inline-flex items-center gap-1.5 text-[10px] font-medium text-primary">
			<span class="relative inline-flex size-1.5">
				<span
					class="absolute inline-flex h-full w-full animate-ping rounded-full bg-primary/50 opacity-75"
				></span>
				<span class="relative inline-flex size-1.5 rounded-full bg-primary"></span>
			</span>
			live
		</span>
	</div>

	<dl class="space-y-1.5">
		{#each rows as row (row.label)}
			<div class="flex items-baseline justify-between text-xs">
				<dt class="text-muted-foreground">{row.label}</dt>
				<dd class="font-mono tabular-nums text-foreground">{row.value}</dd>
			</div>
		{/each}
	</dl>
</div>
