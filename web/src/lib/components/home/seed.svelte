<script lang="ts">
    import { Button, buttonVariants } from "$lib/components/ui/button/index.js";
    import * as Dialog from "$lib/components/ui/dialog/index.js";
    import { createMutation, useQueryClient } from "@tanstack/svelte-query";
    import Spinner from "../ui/spinner/spinner.svelte";
    import { API_URL } from "$lib/queries/api";
    import { toast } from "svelte-sonner";

    let open = $state<boolean>(false);
    const queryClient = useQueryClient();

    const createDeleteMutation = () =>
        createMutation(() => ({
            mutationFn: async () => {
                const res = await fetch(`${API_URL}/seed`, {
                    method: "POST",
                });
                if (!res.ok) throw new Error("Deletion failed");
                return res.json();
            },
            onSuccess: () => {
                queryClient.invalidateQueries({ queryKey: ["getAllData"] });
                toast("seeded successfully");
                open = false;
            },
            onError: () => {
                toast("oops, error");
            },
        }));

    const deltMutation = createDeleteMutation();
</script>

<Dialog.Root bind:open>
    <Dialog.Trigger
        type="button"
        class={buttonVariants({ variant: "outline" })}
    >
        seed
    </Dialog.Trigger>
    <Dialog.Content class="sm:max-w-[425px]">
        <Dialog.Header>
            <Dialog.Title>seed db</Dialog.Title>
        </Dialog.Header>
        <Dialog.Footer
            class="w-full flex flex-col items-center justify-between *:w-full gap-x-4"
        >
            <Button
                disabled={deltMutation.isPending}
                onclick={async () => {
                    await deltMutation.mutateAsync();
                }}
            >
                {#if deltMutation.isPending}
                    <Spinner />
                {/if}
                go ahead</Button
            >
        </Dialog.Footer>
    </Dialog.Content>
</Dialog.Root>
