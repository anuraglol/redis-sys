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

export const GetAllRawSchema = z.object({
  data: z.array(z.string()),
});

export const GetAllTransformSchema = GetAllRawSchema.transform((val) => {
  const result: Record<string, string> = {};
  const array = val.data;

  for (let i = 0; i < array.length; i += 2) {
    const key = array[i];
    const value = array[i + 1];
    if (key !== undefined && value !== undefined) {
      result[key] = value;
    }
  }

  return result;
});

const ValueSchema = z.object({
  value: z.string(),
  expiry: z.coerce.number(),
});

export const GetAllResponseSchema = z.object({
  data: z.record(z.string(), ValueSchema),
});

export type GetAllRes = z.infer<typeof GetAllTransformSchema>;
export type GetRouteInput = z.infer<typeof getRouteSchema>;
export type SetRouteInput = z.infer<typeof setRouteSchema>;
export type DelRouteInput = z.infer<typeof delRouteSchema>;
export type IncrRouteInput = z.infer<typeof incrRouteSchema>;
export type TtlRouteInput = z.infer<typeof ttlRouteSchema>;
export type StatsRouteRes = z.infer<typeof statsRouteSchema>;
