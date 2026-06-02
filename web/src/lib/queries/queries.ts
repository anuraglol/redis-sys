import { createQuery } from "@tanstack/svelte-query";
import { API_URL } from "./api";
import { homeState } from "$lib/state/home.svelte";
import { statsRouteSchema } from "$lib/validations";

export const getStatsQuery = () =>
  createQuery(() => ({
    queryKey: ["stats"],
    queryFn: async () => {
      const res = await fetch(`${API_URL}/stats`);
      if (!res.ok) throw new Error("Failed to fetch server statistics");
      const data: {
        stats: Array<string>;
      } = await res.json();
      const parsed = {};

      if (data && Array.isArray(data.stats)) {
        for (let i = 0; i < data.stats.length; i += 2) {
          parsed[data.stats[i]] = data.stats[i + 1];
        }
      }
      return statsRouteSchema.parse(parsed);
    },
    // refetchInterval: 3000,
  }));

export const getValueQuery = () =>
  createQuery(() => ({
    queryKey: ["value", homeState.selectedKey],
    queryFn: async () => {
      if (!homeState.selectedKey) return null;
      const res = await fetch(`${API_URL}/get/${homeState.selectedKey}`);
      if (res.status === 404) return { value: null, isNil: true };
      if (!res.ok) throw new Error("Failed to retrieve key data");
      return await res.json();
    },
    enabled: !!homeState.selectedKey,
  }));

export const getTtlQuery = () =>
  createQuery(() => ({
    queryKey: ["ttl", homeState.selectedKey],
    queryFn: async () => {
      if (!homeState.selectedKey) return null;
      const res = await fetch(`${API_URL}/ttl/${homeState.selectedKey}`);
      if (!res.ok) return null;
      return await res.json();
    },
    enabled: !!homeState.selectedKey,
  }));

export type StatsQuery = ReturnType<typeof getStatsQuery>;
export type ValueQuery = ReturnType<typeof getValueQuery>;
export type TtlQuery = ReturnType<typeof getTtlQuery>;
