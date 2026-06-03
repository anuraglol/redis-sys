# web - agent guide

sveltekit app, svelte 5 runes, ts strict, tailwind v4, shadcn-svelte (rhea style), bun pkg mgr, tanstack table-core already wired via `src/lib/components/ui/data-table/`.

## build / lint / test

use bun. no test framework configured yet -> add vitest + @testing-library/svelte when first test needed.

- `bun run dev` -> vite dev server
- `bun run build` -> prod build
- `bun run preview` -> preview prod build
- `bun run check` -> svelte-kit sync + svelte-check (ts + svelte typecheck). run after every non-trivial change.
- `bun run check:watch` -> same, watch mode
- `bun run prepare` -> svelte-kit sync (runs on `bun install` via lifecycle)

single test: not configured. once vitest added, use `bunx vitest run path/to/file.test.ts` for one file, `bunx vitest run -t "name"` to filter by test name. add `vitest` + `@testing-library/svelte` + `jsdom` as devDeps, create `vitest.config.ts` extending sveltekit config, add `"test": "vitest"` + `"test:run": "vitest run"` to scripts.

typecheck single file: use svelte-check's `--file` filter if needed, else scope by saving nearby + running `bun run check`. editor intellisense via svelte-vscode ext is the fastest loop.

## project layout

- `src/routes/` -> sveltekit file-based routes (+page.svelte, +layout.svelte, +page.ts, etc)
- `src/lib/` -> `$lib` alias, shared code
- `src/lib/components/ui/` -> shadcn-svelte generated components, one folder per component
- `src/lib/components/ui/<name>/` -> `<name>.svelte` + `index.ts` barrel re-export
- `src/lib/hooks/` -> `.svelte.ts` reactive classes/utilities
- `src/lib/utils.ts` -> `cn()`, `WithElementRef`, `WithoutChild`, `WithoutChildren` types
- `src/lib/assets/` -> static assets importable via `$lib/assets`
- `src/routes/layout.css` -> tailwind v4 entry + css vars (shadcn theme tokens)
- `static/` -> served at root, not importable
- `components.json` -> shadcn-svelte config (aliases, style=rhea, iconLibrary=lucide)

aliases (from components.json + svelte-kit defaults): `$lib`, `$lib/components`, `$lib/components/ui`, `$lib/utils`, `$lib/hooks`.

## code style

### file naming

- kebab-case only: `<name>-<name>.svelte` or `<name>-<name>.ts`. no camelCase, no PascalCase filenames.
- svelte component files match the component: `data-table.svelte`, `flex-render.svelte`.
- reactive logic for a component: same name with `.svelte.ts` suffix (e.g. `data-table.svelte.ts`).
- barrel re-exports: `index.ts` inside the component folder.

### imports

- use `$lib/...` alias, never relative `../../` for lib code.
- svelte 5 imports from `svelte`, `svelte/elements`, `svelte/reactivity` (e.g. `MediaQuery`).
- type-only imports use `type` keyword: `import { type ButtonProps } from "./button.svelte"`.
- barrel pattern: `import { Button } from "$lib/components/ui/button"` re-exports the `Root` as named export.
- shadcn-svelte `cn()` from `$lib/utils` for class merging -> `cn(buttonVariants({ variant, size }), className)`.
- for ui components needing jsx-style attrs, import from `svelte/elements` (`HTMLButtonAttributes`, `HTMLAnchorAttributes`).

### formatting

- tabs for indent (matches existing files).
- double quotes for strings (`"button"`, not `'button'`).
- semicolons on.
- trailing commas in multi-line.
- svelte template attrs: `class={cn(...)}`, `{...restProps}` shorthand, `bind:this={ref}`.
- self-closing tags where possible.

### types

- ts strict mode on (`strict: true` in tsconfig). no `any` unless absolutely needed + eslint-disable comment.
- component props: `type ButtonProps = WithElementRef<HTMLButtonAttributes> & ...` pattern. props go in a `<script lang="ts" module>` block for exports, instance `<script lang="ts">` for `$props()` destructure.
- destructure props with `let { class: className, variant = "default", ref = $bindable(null), ...restProps }: ButtonProps = $props();` pattern.
- export types from `*.svelte` using `<script lang="ts" module>` block.
- `data-slot="button"` style attrs for styling hooks (shadcn convention).

### svelte 5 runes

- runes mode forced via `svelte.config.js` -> use `$props`, `$state`, `$derived`, `$bindable`, `$effect`. no legacy `export let` / `$:`.
- `$props()` for component props, `$bindable()` for two-way bound refs.
- snippets: `{#if href}...{:else}...{/if}` branches, `{@render children?.()}` for slot content.
- reactive classes/utilities in `.svelte.ts` files (not `.ts`) so `$state` works.
- `MediaQuery` from `svelte/reactivity` for reactive media queries (see `is-mobile.svelte.ts`).

### naming

- components: PascalCase (`Button`, `DataTable`).
- variables/fns: camelCase.
- types/interfaces: PascalCase.
- css vars / tailwind tokens: kebab-case (`--background`, `text-foreground`).
- file names: kebab-case only (see above).

### styling

- tailwind v4 utility-first. theme tokens via css vars in `layout.css` (oklch colors).
- shadcn-svelte `rhea` style preset. lucide-svelte for icons (`@lucide/svelte`).
- component variants via `tailwind-variants` (`tv()` -> `buttonVariants`).
- dark mode class on `<html class="dark">` already set in `app.html`. use `dark:` prefix.
- merge classes with `cn()` not template strings.

### error handling

- sveltekit `+page.ts`/`+page.server.ts` load fns throw `error(status, msg)` for http errors, `redirect(status, location)` for redirects.
- async fns in components: handle via `try/catch` in `$effect` or use sveltekit form actions + superforms (`sveltekit-superforms` already installed).
- toast errors via `svelte-sonner` (`toast.error(...)`).

## general rules

- **lowercase only**: all comms lowercase. except code blocks + verbatim quoted content.
- **caveman mode**: drop articles (a/an/the), filler (just/really/basically/actually/simply), pleasantries (sure/certainly/of course/happy to), hedging. fragments ok. short synonyms (big not extensive, fix not "implement a solution for"). abbreviate common terms (db/auth/config/req/res/fn/impl). strip conjunctions. use arrows for causality (X -> Y). one word when one word enough. technical terms stay exact. code blocks unchanged. errors quoted exact.
- **no comments**: do not add any comments to code. ever. even helpful ones. even "obvious" ones. if you feel urge to comment -> refactor code to be self-explanatory, or use better naming.
- **no proactive commits**: never `git commit` unless user explicitly asks.
- **no new files unless required**: edit existing files when possible.

## data layer

- @tanstack/table-core already installed, wired via `createSvelteTable` in `src/lib/components/ui/data-table/data-table.svelte.ts`. use that for tables.
- **add @tanstack/svelte-query** for server state when needed. install: `bun add @tanstack/svelte-query`. wrap root `+layout.svelte` with `QueryClientProvider` + `setQueryClientContext`. use `createQuery`, `createMutation` from `@tanstack/svelte-query`. default staleTime 30s, refetchOnWindowFocus false unless explicitly needed.

## patterns

- **table**: `const table = createSvelteTable({ data, columns, getCoreRowModel: getCoreRowModel() })`. render with `FlexRender` for header/cell content (snippets or components). use `renderComponent` / `renderSnippet` from `data-table/render-helpers` to bridge column defs to svelte.
- **icons**: `import { IconName } from "@lucide/svelte"`. size via class (`size-4`, `size-5`) not the `size` prop when possible.
- **media query hook**: extend `MediaQuery` from `svelte/reactivity` (see `is-mobile.svelte.ts`) instead of hand-rolling matchMedia + state.
- **forms**: `sveltekit-superforms` + `formsnap` + `zod` (or arktype/valibot) -> action in `+page.server.ts`, schema validated server-side, client uses `superForm` + `Form` from `formsnap`.
- **toasts**: `import { toast } from "svelte-sonner"` then `toast.success(msg)` / `toast.error(msg)`. mount `<Toaster />` once in root `+layout.svelte`.
- **theme**: `mode-watcher` handles dark/light. use `import { mode } from "mode-watcher"` for reactive theme access.

## gotchas

- `<script lang="ts" module>` block is hoisted (runs once at module load) -> use for exports + `tv()` calls + types. `<script lang="ts">` is per-instance -> use for `$props()` destructure.
- `$lib` imports use `.js` extension when importing from `.svelte` files in ts config: `import { cn } from "$lib/utils.js"` (the bundler resolves it; ts needs the explicit ext).
- runes mode forced globally -> never use `export let`, `$:`, `on:click` (use `onclick` props instead), `<slot>` (use `{@render children?.()}`).
- shadcn `rhea` style is the active preset. when adding new ui components via `bunx shadcn-svelte@latest add <name>`, output goes to `$lib/components/ui/<name>/` and may overwrite styles -> review diff.
- `tsconfig` extends generated `.svelte-kit/tsconfig.json` -> don't edit generated files. rerun `bun run prepare` after `svelte.config.js` changes.
- dark mode is always-on (`<html class="dark">` in `app.html`). light mode toggle via `mode-watcher` removes/adds the class.

## file creation rules

- file names: kebab-case, `<name>-<name>.svelte` or `<name>-<name>.ts`. dashes between words. no underscores, no camelCase in filenames.
- match existing extension to content: `.svelte` for components, `.ts` for pure logic, `.svelte.ts` for logic needing runes, `.css` for styles.

## key deps (do not re-add)

svelte 5, sveltekit 2, @sveltejs/adapter-auto, bits-ui, formsnap, layerchart, mode-watcher, paneforge, shadcn-svelte, svelte-sonner, sveltekit-superforms, tailwindcss v4 + @tailwindcss/vite, tailwind-variants, tw-animate-css, embla-carousel-svelte, vaul-svelte, @lucide/svelte, @internationalized/date, @fontsource-variable/inter, clsx, tailwind-merge.
