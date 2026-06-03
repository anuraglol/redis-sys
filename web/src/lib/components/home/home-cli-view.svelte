<script lang="ts">
	import { ArrowRight, Terminal } from "@lucide/svelte";
	import Button from "$lib/components/ui/button/button.svelte";
	import { executeCliCommand, type CliOutputLine } from "$lib/queries/cli";

	let cliCommand = $state("");
	let cliOutput = $state<CliOutputLine[]>([]);

	async function handleSubmit(e: SubmitEvent) {
		e.preventDefault();
		if (!cliCommand.trim()) return;
		const line = await executeCliCommand(cliCommand);
		cliOutput = [...cliOutput, line];
		cliCommand = "";
	}
</script>

<header
	class="h-16 border-b border-border px-8 flex items-center justify-between bg-card/30 backdrop-blur-md shrink-0"
>
	<div class="flex items-center gap-3">
		<Terminal class="w-5 h-5 text-primary" />
		<h2 class="text-sm font-semibold tracking-tight">
			Interactive Query Pipeline Shell
		</h2>
	</div>
	<Button
		variant="outline"
		size="sm"
		class="h-8 text-xs font-mono text-muted-foreground"
		onclick={() => (cliOutput = [])}
	>
		Flush Shell Output
	</Button>
</header>

<div
	class="flex-1 bg-black text-zinc-200 font-mono p-6 overflow-y-auto flex flex-col justify-between selection:bg-zinc-800"
>
	<div class="space-y-3 flex-1 overflow-y-auto mb-4 pr-2">
		<div class="text-zinc-500 text-xs border-b border-zinc-900 pb-3">
			<div>
				DrizzleRedis Reactive Cluster Command Shell Terminal Console [Version
				2026.1]
			</div>
			<div>
				Enter native engine mutations: SET key value, GET key, DEL key, INCR
				key, TTL key, STATS
			</div>
		</div>

		{#each cliOutput as line (line.cmd + line.raw)}
			<div class="space-y-1 text-xs">
				<div class="flex items-center gap-2 text-zinc-400">
					<ArrowRight class="w-3 h-3 text-primary animate-pulse" />
					<span>{line.cmd}</span>
				</div>
				<div
					class="pl-5 whitespace-pre-wrap {line.error
						? 'text-rose-500 font-bold'
						: 'text-emerald-400'}"
				>
					{line.raw}
				</div>
			</div>
		{/each}
	</div>

	<form
		onsubmit={handleSubmit}
		class="flex items-center gap-3 bg-zinc-950 p-2.5 rounded-xl border border-zinc-900 shadow-2xl shrink-0"
	>
		<span class="text-primary font-bold text-sm pl-2"
			>redis://127.0.0.1:8001></span
		>
		<input
			type="text"
			placeholder="SET user:name 'Go' EX 3600"
			bind:value={cliCommand}
			class="flex-1 bg-transparent border-0 outline-none focus:ring-0 text-xs text-white placeholder:text-zinc-700 font-mono"
		/>
		<Button
			type="submit"
			size="sm"
			class="h-7 px-3 text-[10px] uppercase font-bold bg-white text-black hover:bg-zinc-200 rounded-lg"
		>
			Send
		</Button>
	</form>
</div>
