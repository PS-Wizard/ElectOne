<script lang="ts">
    import { goto } from "$app/navigation";

    let email = "";
    let password = "";
    let totp_code = "";
    let qr_url = "";
    let setup_done = true;
    let token = "";
    let error = "";

    const login = async () => {
        error = "";

        try {
            const res = await fetch("http://localhost:3000/auth/admin", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ email, password, totp_code }),
            });

            const text = await res.text(); // safer than res.json()

            const data = (() => {
                try {
                    return JSON.parse(text);
                } catch {
                    return { message: text }; // fallback to plain text
                }
            })();

            if (!res.ok) {
                error = data.message || "Login failed";
                return;
            }

            if (!data.setup_done && data.qr_url) {
                setup_done = false;
                qr_url = `https://api.qrserver.com/v1/create-qr-code/?size=200x200&data=${encodeURIComponent(data.qr_url)}`;
            } else {
                token = data.token;
                localStorage.setItem("admin_token", token);
                goto("/admin/appeals");
            }
        } catch (err) {
            error = "Something went wrong. Please try again.";
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
            <img
                src="https://merakiui.com/images/logo.svg"
                class="mx-auto h-8"
                alt="Logo"
            />
            <h2 class="mt-4 text-2xl font-bold text-gray-800 dark:text-white">
                Admin Login
            </h2>
            <p class="text-sm text-gray-500 dark:text-gray-400">
                Access the control panel
            </p>
        </div>

        <div class="space-y-4">
            <input
                type="email"
                placeholder="Email"
                class="input input-bordered w-full"
                bind:value={email}
                required
            />
            <input
                type="password"
                placeholder="Password"
                class="input input-bordered w-full"
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
                    class="input input-bordered w-full"
                    bind:value={totp_code}
                />
                <p
                    class="text-xs text-gray-100 text-right dark:text-gray-400 mt-1"
                >
                    If this is your first time, leave this blank.
                </p>
            </div>

            {#if error}
                <div class="alert alert-error">{error}</div>
            {/if}

            <button
                class="btn btn-primary w-full"
                on:click|preventDefault={login}
            >
                Login
            </button>

            {#if token}
                <div class="alert alert-success">
                    Logged in! Token: <br /><code class="break-all"
                        >{token}</code
                    >
                </div>
            {/if}
        </div>
    </div>
</section>
