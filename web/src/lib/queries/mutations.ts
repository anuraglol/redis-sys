import { createMutation } from "@tanstack/svelte-query";
import { API_URL } from "./api";
import { queryClient } from "./query-client";
import { homeState } from "$lib/state/home.svelte";

export const createSetMutation = () =>
	createMutation(() => ({
		mutationFn: async (payload) => {
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
			queryClient.invalidateQueries({ queryKey: ["stats"] });
			if (homeState.formKey === homeState.selectedKey) {
				queryClient.invalidateQueries({
					queryKey: ["value", homeState.selectedKey],
				});
				queryClient.invalidateQueries({
					queryKey: ["ttl", homeState.selectedKey],
				});
			}
			homeState.resetForm();
		},
	}));

export const createDeleteMutation = () =>
	createMutation(() => ({
		mutationFn: async (keyStr: string) => {
			const res = await fetch(`${API_URL}/del/${keyStr}`, {
				method: "DELETE",
			});
			if (!res.ok) throw new Error("Deletion failed");
			return res.json();
		},
		onSuccess: (_, deletedKeyStr: string) => {
			queryClient.invalidateQueries({ queryKey: ["stats"] });
			if (homeState.selectedKey === deletedKeyStr) {
				homeState.selectedKey = "";
			}
		},
	}));

export const createIncrMutation = () =>
	createMutation(() => ({
		mutationFn: async (keyStr: string) => {
			const res = await fetch(`${API_URL}/incr/${keyStr}`, {
				method: "POST",
			});
			if (!res.ok) throw new Error("Increment operations failed");
			return res.json();
		},
		onSuccess: (_, targetKey: string) => {
			if (homeState.selectedKey === targetKey) {
				queryClient.invalidateQueries({
					queryKey: ["value", homeState.selectedKey],
				});
			}
		},
	}));
