import { z } from "zod";

const keyParamSchema = z
  .string()
  .min(1, "Key is required")
  .refine((val) => !val.includes("/"), {
    message: "Key cannot contain slashes",
  });

export const getRouteSchema = z.object({
  params: z.object({
    key: keyParamSchema,
  }),
});

export const setRouteSchema = z.object({
  body: z.object({
    key: z.string().min(1, "Key is required"),
    value: z.string().min(1, "Value is required"),
    EX: z.number().int().positive().optional(),
  }),
});

export const delRouteSchema = z.object({
  params: z.object({
    key: keyParamSchema,
  }),
});

export const incrRouteSchema = z.object({
  params: z.object({
    key: keyParamSchema,
  }),
});

export const ttlRouteSchema = z.object({
  params: z.object({
    key: keyParamSchema,
  }),
});

export const statsRouteSchema = z.object({
  total_commands: z.string(),
  set_cmds: z.string(),
  get_cmds: z.string(),
  ping_cmds: z.string(),
  expire_cmds: z.string(),
});

export type GetRouteInput = z.infer<typeof getRouteSchema>;
export type SetRouteInput = z.infer<typeof setRouteSchema>;
export type DelRouteInput = z.infer<typeof delRouteSchema>;
export type IncrRouteInput = z.infer<typeof incrRouteSchema>;
export type TtlRouteInput = z.infer<typeof ttlRouteSchema>;
export type StatsRouteRes = z.infer<typeof statsRouteSchema>;
