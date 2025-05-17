<script lang="ts">
	import Navbar from "../../../components/Navbar.svelte";

	let citizenship_id = "";
	let voter_card_id = "";
	let password = "";
	let totp = "";
	let error = "";
	let success = "";

	function isValidCitizenship(id: string) {
		return /^\d{2}-\d{2}-\d{2}-\d{5}$/.test(id);
	}

	function isValidVoterCard(id: string) {
		return /^\d{10}$/.test(id);
	}

	const deleteAccount = async () => {
		error = "";
		success = "";

		if (!isValidCitizenship(citizenship_id)) {
			error = "Invalid Citizenship ID format.";
			return;
		}

		if (!isValidVoterCard(voter_card_id)) {
			error = "Invalid Voter Card ID.";
			return;
		}

		if (password.length < 8) {
			error = "Password seems too short.";
			return;
		}

		try {
			const res = await fetch("http://localhost:3000/auth/delete", {
				method: "POST",
				headers: { "Content-Type": "application/json" },
				body: JSON.stringify({
					citizenship_id,
					voter_card_id,
					password,
					totp,
				}),
			});

			const text = await res.text();
			const data = (() => {
				try {
					return JSON.parse(text);
				} catch {
					return { message: text };
				}
			})();

			if (!res.ok) {
				error = data.error || data.message || "Delete failed.";
				return;
			}

			alert("Account deleted. Sad to see you go.");
            localStorage.removeItem("admin_token");
		} catch (err) {
			error = "Something went wrong.";
			console.error(err);
		}
	};
</script>

<Navbar />

<div class="min-h-screen flex items-center justify-center bg-transparent p-6">
	<div class="w-full max-w-md bg-white rounded-xl shadow-lg p-10">
		<div class="text-left mb-8">
			<h1 class="text-3xl font-bold text-black">Delete Account</h1>
			<p class="text-gray-600 text-sm mt-2">
				Confirm your identity to delete your account. This cannot be undone.
			</p>
		</div>

		<form on:submit|preventDefault={deleteAccount} class="space-y-6">
			<div>
				<label class="block text-sm font-medium text-gray-700 mb-1">Citizenship ID</label>
				<input
					type="text"
					bind:value={citizenship_id}
					required
					placeholder="XX-XX-XX-XXXXX"
					class="input w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black"
				/>
			</div>

			<div>
				<label class="block text-sm font-medium text-gray-700 mb-1">Voter Card ID</label>
				<input
					type="text"
					bind:value={voter_card_id}
					required
					placeholder="XXXXXXXXXX"
					class="input w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black"
				/>
			</div>

			<div>
				<label class="block text-sm font-medium text-gray-700 mb-1">Password</label>
				<input
					type="password"
					bind:value={password}
					required
					placeholder="Your Password"
					class="input w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black"
				/>
			</div>

			<div>
				<label class="block text-sm font-medium text-gray-700 mb-1">TOTP Code</label>
				<input
					type="text"
					bind:value={totp}
					required
					placeholder="TOTP Code"
					class="input w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black"
				/>
			</div>

			{#if error}
				<div class="text-sm text-red-500 text-center font-medium">
					{error}
				</div>
			{/if}

			{#if success}
				<div class="text-sm text-green-600 text-center font-medium">
					{success}
				</div>
			{/if}

			<button
				type="submit"
				class="btn w-full uppercase tracking-wide font-bold bg-red-600 hover:bg-red-700 text-white transition duration-300"
			>
				Delete My Account
			</button>
		</form>
	</div>
</div>
