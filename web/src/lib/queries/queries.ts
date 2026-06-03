import { createQuery } from "@tanstack/svelte-query";
import { API_URL } from "./api";
import { homeState } from "$lib/state/home.svelte";
import { GetAllResponseSchema, GetAllTransformSchema, statsRouteSchema } from "$lib/validations";

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

export const getAllQuery = () =>
  createQuery(() => ({
    queryKey: ["getAllData"],
    queryFn: async () => {
      const response = await fetch(`${API_URL}/getall`);
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }

      return GetAllResponseSchema.parse(await response.json());
    },
  }));

export type StatsQuery = ReturnType<typeof getStatsQuery>;
