<script lang="ts">
    import { Button, buttonVariants } from "$lib/components/ui/button/index.js";
    import * as Dialog from "$lib/components/ui/dialog/index.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import { Label } from "$lib/components/ui/label/index.js";
    import { API_URL } from "$lib/queries/api";
    import { createMutation, useQueryClient } from "@tanstack/svelte-query";

    const queryClient = useQueryClient();

    let open = $state(false);
    let key = $state("");
    let formValue = $state("");
    let formTtl = $state("");

    const createSetMutation = () =>
        createMutation(() => ({
            mutationFn: async (payload: {
                key: string;
                value: string;
                ttl: string;
            }) => {
                const body: { key: string; value: string; ex?: number } = {
                    key: payload.key,
                    value: payload.value,
                };
                if (payload.ttl) body.ex = parseInt(payload.ttl, 10);
                const res = await fetch(`${API_URL}/set`, {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify(body),
                });
                if (!res.ok) throw new Error("Failed to write token");
                return res.json();
            },
            onSuccess: () => {
                queryClient.invalidateQueries({ queryKey: ["getAllData"] });
                key = "";
                formValue = "";
                formTtl = "";
                open = false;
            },
            onError: (e) => {
                console.error(e);
            },
        }));

    const setMutation = createSetMutation();

    async function handleSubmit(e: SubmitEvent) {
        e.preventDefault();
        await setMutation.mutateAsync({
            key,
            value: formValue,
            ttl: formTtl,
        });
    }
</script>

<Dialog.Root bind:open>
    <Dialog.Trigger
        type="button"
        class={buttonVariants({ variant: "outline" })}
    >
        add entry
    </Dialog.Trigger>
    <Dialog.Content class="sm:max-w-[425px]">
        <Dialog.Header>
            <Dialog.Title>add entry</Dialog.Title>
        </Dialog.Header>

        <form onsubmit={handleSubmit} class="grid gap-4">
            <div class="grid gap-3">
                <Label for="form-key">key</Label>
                <Input id="form-key" required bind:value={key} />
            </div>

            <div class="grid gap-3">
                <Label for="form-value">value</Label>
                <Input id="form-value" required bind:value={formValue} />
            </div>

            <div class="grid gap-3">
                <Label for="form-ttl">expiry (seconds)</Label>
                <Input id="form-ttl" bind:value={formTtl} />
            </div>

            <Dialog.Footer class="mt-2">
                <Dialog.Close
                    type="button"
                    class={buttonVariants({ variant: "outline" })}
                >
                    Cancel
                </Dialog.Close>
                <Button type="submit" disabled={setMutation.isPending}>
                    {setMutation.isPending ? "Creating..." : "Create entry"}
                </Button>
            </Dialog.Footer>
        </form>
    </Dialog.Content>
</Dialog.Root>
