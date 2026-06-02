import { API_URL } from "./api";
import { queryClient } from "./query-client";

export type CliOutputLine = {
	cmd: string;
	raw: string;
	error: boolean;
};

export async function executeCliCommand(command: string): Promise<CliOutputLine> {
	const line: CliOutputLine = { cmd: command, raw: "", error: false };
	const parts = command.trim().split(/\s+/);
	const cmd = parts[0].toUpperCase();
	const args = parts.slice(1);

	try {
		if (cmd === "SET") {
			const body: { key: string; value: string; ex?: number } = {
				key: args[0],
				value: args[1],
			};
			if (args[2]?.toUpperCase() === "EX" && args[3])
				body.ex = parseInt(args[3], 10);
			const res = await fetch(`${API_URL}/set`, {
				method: "POST",
				headers: { "Content-Type": "application/json" },
				body: JSON.stringify(body),
			});
			line.raw = await res.text();
		} else if (cmd === "GET") {
			const res = await fetch(`${API_URL}/get/${args[0]}`);
			line.raw = await res.text();
		} else if (cmd === "DEL") {
			const res = await fetch(`${API_URL}/del/${args[0]}`, {
				method: "DELETE",
			});
			line.raw = await res.text();
		} else if (cmd === "INCR") {
			const res = await fetch(`${API_URL}/incr/${args[0]}`, {
				method: "POST",
			});
			line.raw = await res.text();
		} else if (cmd === "TTL") {
			const res = await fetch(`${API_URL}/ttl/${args[0]}`);
			line.raw = await res.text();
		} else if (cmd === "STATS") {
			const res = await fetch(`${API_URL}/stats`);
			line.raw = await res.text();
		} else {
			throw new Error(`ERR unknown command '${cmd}'`);
		}
		queryClient.invalidateQueries({ queryKey: ["stats"] });
	} catch (err) {
		line.error = true;
		line.raw = err instanceof Error ? err.message : String(err);
	}

	return line;
}
