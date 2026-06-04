<script lang="ts">
	import { page } from "$app/state";
	import { Database, Terminal } from "@lucide/svelte";

	const items: Array<{ href: string; label: string; icon: typeof Database }> = [
		{ href: "/", label: "entries", icon: Database },
		{ href: "/cli", label: "shell", icon: Terminal },
	];
</script>

<nav class="flex flex-col gap-0.5">
	{#each items as item (item.href)}
		{@const Icon = item.icon}
		{@const active = page.url.pathname === item.href}
		<a
			href={item.href}
			class="group relative flex items-center gap-2.5 rounded-xl px-3 py-2 text-sm transition-all duration-150 outline-none focus-visible:ring-2 focus-visible:ring-ring/40
				{active
				? 'bg-card text-foreground shadow-[0_1px_0_0_color-mix(in_oklch,var(--primary),transparent_85%)] ring-1 ring-sidebar-border'
				: 'text-muted-foreground hover:text-foreground hover:bg-sidebar-accent/60'}"
			aria-current={active ? "page" : undefined}
		>
			<Icon class="size-4 {active ? 'text-primary' : ''}" />
			<span class="font-medium">{item.label}</span>
			{#if active}
				<span class="ml-auto inline-flex size-1.5 rounded-full bg-primary"></span>
			{/if}
		</a>
	{/each}
</nav>
