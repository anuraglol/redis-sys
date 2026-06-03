<script lang="ts">
	import { Database } from "@lucide/svelte";
	import KeyAnalyzerCard from "./key-analyzer-card.svelte";
	import KeyMutatorCard from "./key-mutator-card.svelte";
</script>

<header
	class="h-16 border-b border-border px-8 flex items-center justify-between bg-card/30 backdrop-blur-md shrink-0"
>
	<div class="flex items-center gap-3">
		<Database class="w-5 h-5 text-muted-foreground" />
		<h2 class="text-sm font-semibold tracking-tight">
			Active Schema Memory Engine
		</h2>
	</div>
</header>

<div class="flex-1 overflow-y-auto p-8 space-y-8">
	<div class="grid grid-cols-1 xl:grid-cols-3 gap-8 items-start">
		<KeyAnalyzerCard />
		<KeyMutatorCard />
	</div>
</div>
