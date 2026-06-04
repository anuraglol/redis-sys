<script lang="ts">
	import SidebarTabs from "./sidebar-tabs.svelte";
	import SidebarStatsFooter from "./sidebar-stats-footer.svelte";
	import { getStatsQuery } from "$lib/queries/queries";

	const statsQuery = getStatsQuery();
</script>

<aside
	class="w-64 shrink-0 border-r border-sidebar-border bg-sidebar text-sidebar-foreground flex flex-col z-10"
>
	<div class="px-5 pt-6 pb-5">
		<a href="/" class="flex items-center gap-2.5 outline-none focus-visible:ring-2 focus-visible:ring-ring/40 rounded-md">
			<span class="inline-flex items-center justify-center text-primary animate-bob">
				<svg viewBox="0 0 16 16" fill="none" class="size-4" aria-hidden="true">
					<path
						d="M8 0.8 L9.35 6.65 L15.2 8 L9.35 9.35 L8 15.2 L6.65 9.35 L0.8 8 L6.65 6.65 Z"
						fill="currentColor"
					/>
				</svg>
			</span>
			<h1 class="font-serif text-2xl leading-none tracking-tight">kotiri</h1>
		</a>
		<p class="mt-2 text-xs text-muted-foreground pl-[26px]">
			tiny in-memory store
		</p>
	</div>

	<div class="px-3">
		<SidebarTabs />
	</div>

	<div class="flex-1"></div>

	<SidebarStatsFooter {statsQuery} />
</aside>
