<script lang="ts">
	import { AlertTriangle, KeyRound } from "@lucide/svelte";
	import Badge from "$lib/components/ui/badge/badge.svelte";
	import { homeState } from "$lib/state/home.svelte";
	import type { StatsQuery } from "$lib/queries/queries";

	let {
		statsQuery,
		onTabChange,
	}: {
		statsQuery: StatsQuery;
		onTabChange: (tab: string) => void;
	} = $props();

	let keysList = $derived.by(() => {
		if (!statsQuery.data) return [];
		return Object.keys(statsQuery.data)
			.filter((k) => k.endsWith("_cmds"))
			.map((k) => k.replace("_cmds", ""));
	});
</script>

<div>
	<div
		class="text-xs font-semibold text-muted-foreground uppercase tracking-wider mb-2 px-2 flex items-center justify-between"
	>
		<span>Tracked Operational Triggers</span>
		<Badge variant="outline" class="font-mono text-[10px] px-1 py-0"
			>{keysList.length}</Badge
		>
	</div>
	{#if statsQuery.isLoading}
		<div class="space-y-1.5 px-2">
			<div class="h-8 bg-muted rounded-lg animate-pulse w-full"></div>
			<div class="h-8 bg-muted rounded-lg animate-pulse w-full"></div>
			<div class="h-8 bg-muted rounded-lg animate-pulse w-full"></div>
		</div>
	{:else if statsQuery.error}
		<div
			class="p-3 text-xs bg-destructive/10 text-destructive rounded-lg border border-destructive/20 flex gap-2"
		>
			<AlertTriangle class="h-4 w-4 shrink-0" />
			<span>Offline Node Cluster Connections</span>
		</div>
	{:else}
		<div class="space-y-0.5">
			{#each keysList as engineKey}
				<button
					onclick={() => {
						homeState.selectedKey = engineKey;
						onTabChange("data");
					}}
					class="w-full text-left px-3 py-2 rounded-lg text-sm transition-all flex items-center justify-between group {homeState.selectedKey ===
					engineKey
						? 'bg-primary text-primary-foreground font-medium shadow-sm'
						: 'hover:bg-muted text-foreground/80 hover:text-foreground'}"
				>
					<div class="flex items-center gap-2.5 truncate">
						<KeyRound
							class="w-4 h-4 shrink-0 opacity-70 {homeState.selectedKey ===
							engineKey
								? 'text-primary-foreground'
								: 'text-primary'}"
						/>
						<span class="truncate font-mono text-xs">{engineKey}</span>
					</div>
					<Badge
						variant="secondary"
						class="font-mono text-[10px] tracking-tighter shrink-0 {homeState.selectedKey ===
						engineKey
							? 'bg-primary-foreground/20 text-primary-foreground'
							: 'bg-muted-foreground/10 text-muted-foreground'}"
					>
						{statsQuery.data[engineKey + "_cmds"] || 0} ops
					</Badge>
				</button>
			{/each}
			{#if keysList.length === 0}
				<div
					class="text-center py-8 text-xs text-muted-foreground border border-dashed rounded-xl border-border"
				>
					No targeted system keys active.
				</div>
			{/if}
		</div>
	{/if}
</div>
