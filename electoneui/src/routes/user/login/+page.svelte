<script lang="ts">
    import { goto } from "$app/navigation";

    let citizenship_id = "";
    let voter_card_id = "";
    let password = "";
    let totp_code = "";
    let qr_url = "";
    let setup_done = true;
    let token = "";
    let error = "";

    const login = async () => {
        error = "";

        try {
            const res = await fetch("http://localhost:3000/auth/user", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    citizenship_id,
                    voter_card_id,
                    password,
                    totp_code,
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
                error = data.error || data.message || "Login failed";
                return;
            }

            if (data.qr_url) {
                setup_done = false;
                qr_url = `https://api.qrserver.com/v1/create-qr-code/?size=200x200&data=${encodeURIComponent(data.qr_url)}`;
                return;
            }

            token = data.token;
            localStorage.setItem("token", token);
            goto("/user/dashboard");
        } catch (err) {
            error = "Something went wrong.";
            console.error(err);
        }
    };
</script>

<section
    class="min-h-screen flex items-center justify-center bg-gray-100 dark:bg-gray-900"
>
    <div
        class="w-full max-w-md p-8 space-y-6 bg-white rounded-xl shadow-lg dark:bg-gray-800"
    >
        <div class="text-center">
            <h2 class="text-2xl font-bold text-gray-800 dark:text-white">
                User Login
            </h2>
            <p class="text-sm text-gray-500 dark:text-gray-400">
                Login to vote securely
            </p>
        </div>

        <div class="space-y-4">
            <input
                type="text"
                placeholder="Citizenship ID"
                class="input rounded-lg input-bordered w-full"
                bind:value={citizenship_id}
                required
            />
            <input
                type="text"
                placeholder="Voter Card ID"
                class="input input-bordered w-full rounded-lg"
                bind:value={voter_card_id}
                required
            />
            <input
                type="password"
                placeholder="Password"
                class="input input-bordered w-full rounded-lg"
                bind:value={password}
                required
            />

            {#if !setup_done}
                <div class="text-sm text-yellow-500">
                    Scan this QR code with your authenticator app:
                </div>
                <img
                    src={qr_url}
                    alt="QR Code"
                    class="mx-auto my-2 rounded-md shadow-md"
                />
            {/if}

            <div>
                <input
                    type="text"
                    placeholder="TOTP Code"
                    class="rounded-lg input input-bordered w-full"
                    bind:value={totp_code}
                />
                <p class="text-xs text-gray-300 text-right mt-1">
                    If this is your first login, leave this blank.
                </p>
            </div>

            {#if error}
                <div class="alert alert-error">{error}</div>
            {/if}

            <button
                class="rounded-lg btn btn-primary w-full"
                on:click|preventDefault={login}
            >
                Login
            </button>

            {#if token}
                <div class="alert alert-success">
                    Logged in! Token:<br /><code class="break-all">{token}</code
                    >
                </div>
            {/if}
        </div>
    </div>
</section>
