<script lang="ts">
	import { Check, Play, RefreshCw, ShieldAlert } from "@lucide/svelte";
	import * as Card from "$lib/components/ui/card";
	import * as Alert from "$lib/components/ui/alert";
	import Button from "$lib/components/ui/button/button.svelte";
	import Input from "$lib/components/ui/input/input.svelte";
	import { homeState } from "$lib/state/home.svelte";
	import { createSetMutation } from "$lib/queries/mutations";

	const setMutation = createSetMutation();
</script>

<Card.Root class="shadow-sm border border-border bg-card">
	<Card.Header class="pb-3 border-b border-border/60">
		<Card.Title class="text-base font-bold tracking-tight"
			>Data Mutator Engine</Card.Title
		>
		<Card.Description class="text-xs"
			>Inject key structures directly inside the reactive runtime layer</Card.Description
		>
	</Card.Header>
	<Card.Content class="pt-6">
		<form
			onsubmit={(e) => {
				e.preventDefault();
				if (homeState.formKey && homeState.formValue)
					setMutation.mutate({
						key: homeState.formKey,
						value: homeState.formValue,
						ttl: homeState.formTtl,
					});
			}}
			class="space-y-4"
		>
			<div class="space-y-1.5">
				<label
					for="formKey"
					class="text-xs font-semibold text-muted-foreground font-mono"
					>Identifier Key</label
				>
				<Input
					id="formKey"
					type="text"
					placeholder="e.g. user:session:100"
					bind:value={homeState.formKey}
					class="h-9 font-mono text-xs"
					required
				/>
			</div>
			<div class="space-y-1.5">
				<label
					for="formValue"
					class="text-xs font-semibold text-muted-foreground font-mono"
					>Value Payload Data</label
				>
				<Input
					id="formValue"
					type="text"
					placeholder="String data injection payload"
					bind:value={homeState.formValue}
					class="h-9 font-mono text-xs"
					required
				/>
			</div>
			<div class="space-y-1.5">
				<label
					for="formTtl"
					class="text-xs font-semibold text-muted-foreground font-mono"
					>Expiration Window (EX Secs)</label
				>
				<Input
					id="formTtl"
					type="number"
					placeholder="Optional persistence bounds"
					bind:value={homeState.formTtl}
					class="h-9 font-mono text-xs"
				/>
			</div>

			{#if setMutation.isError}
				<Alert.Root variant="destructive" class="py-2.5 px-3">
					<ShieldAlert class="h-4 w-4" />
					<Alert.Description class="text-xs font-mono"
						>{setMutation.error.message}</Alert.Description
					>
				</Alert.Root>
			{/if}

			{#if setMutation.isSuccess}
				<div
					class="p-2.5 rounded-lg bg-emerald-500/10 border border-emerald-500/20 text-emerald-500 text-xs font-mono flex items-center gap-2"
				>
					<Check class="w-4 h-4" /> Operation committed successfully.
				</div>
			{/if}

			<Button
				type="submit"
				class="w-full h-9 text-xs font-semibold tracking-wide"
				disabled={setMutation.isPending}
			>
				{#if setMutation.isPending}
					<RefreshCw class="w-3.5 h-3.5 mr-2 animate-spin" /> Committing Stream...
				{:else}
					<Play class="w-3.5 h-3.5 mr-2 fill-current" /> Execute Mutator Transaction
				{/if}
			</Button>
		</form>
	</Card.Content>
</Card.Root>
