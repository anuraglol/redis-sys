<script lang="ts">
	import { ChevronRight, Eraser, Terminal } from "@lucide/svelte";
	import { Button } from "$lib/components/ui/button";
	import { executeCliCommand, type CliOutputLine } from "$lib/queries/cli";
	import { tick } from "svelte";

	let cliCommand = $state("");
	let cliOutput = $state<CliOutputLine[]>([]);
	let scrollRef = $state<HTMLDivElement | null>(null);

	async function handleSubmit(e: SubmitEvent) {
		e.preventDefault();
		if (!cliCommand.trim()) return;
		const line = await executeCliCommand(cliCommand);
		cliOutput = [...cliOutput, line];
		cliCommand = "";
		await tick();
		scrollRef?.scrollTo({ top: scrollRef.scrollHeight, behavior: "smooth" });
	}
</script>

<header
	class="h-14 border-b border-border px-6 flex items-center justify-between bg-background/80 backdrop-blur shrink-0"
>
	<div class="flex items-center gap-2.5">
		<Terminal class="size-4 text-primary" />
		<h2 class="font-serif text-lg leading-none">shell</h2>
		<span class="text-xs text-muted-foreground ml-2">type a command, hit enter</span>
	</div>
	<Button
		variant="ghost"
		size="sm"
		class="h-7 gap-1.5 text-xs text-muted-foreground"
		onclick={() => (cliOutput = [])}
	>
		<Eraser class="size-3.5" />
		clear
	</Button>
</header>

<div class="flex-1 overflow-hidden p-6">
	<label
		for="cli-input"
		class="block w-full h-full rounded-2xl bg-terminal-bg text-terminal-fg border border-terminal-border shadow-sm overflow-hidden cursor-text"
	>
		<div class="flex flex-col h-full">
			<div bind:this={scrollRef} class="flex-1 overflow-y-auto px-5 py-4 font-mono text-xs leading-relaxed">
				<div class="text-terminal-muted pb-3 mb-3 border-b border-terminal-border space-y-0.5">
					<div>welcome to kotiri shell</div>
					<div>commands: <span class="text-terminal-fg">SET GET DEL INCR TTL STATS</span></div>
				</div>

				{#each cliOutput as line, i (i)}
					<div class="mb-3 space-y-1">
						<div class="flex items-center gap-1.5 text-terminal-prompt">
							<ChevronRight class="size-3" />
							<span>{line.cmd}</span>
						</div>
						<div
							class="pl-4 whitespace-pre-wrap break-words {line.error
								? 'text-terminal-err'
								: 'text-terminal-ok'}"
						>
							{line.raw || "(empty)"}
						</div>
					</div>
				{/each}
			</div>

			<form
				onsubmit={handleSubmit}
				class="flex items-center gap-2 px-4 py-3 border-t border-terminal-border bg-black/15"
			>
				<span class="font-mono text-xs text-terminal-prompt select-none">kotiri</span>
				<span class="font-mono text-xs text-terminal-muted select-none">›</span>
				<input
					id="cli-input"
					type="text"
					placeholder='SET hello "world"'
					bind:value={cliCommand}
					autocomplete="off"
					autocorrect="off"
					autocapitalize="off"
					spellcheck="false"
					class="flex-1 bg-transparent border-0 outline-none text-xs text-terminal-fg placeholder:text-terminal-muted font-mono"
				/>
			</form>
		</div>
	</label>
</div>
