<script lang="ts">
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import Navbar from "../../../components/Navbar.svelte";

    let citizenship_id = "";
    let voter_card_id = "";
    let password = "";
    let totp_code = "";
    let qr_url = "";
    let setup_done = true;
    let token = "";
    let error = "";

    onMount(() => {
        const token = localStorage.getItem("user_token");
        if (token) {
            goto("/user/dashboard");
        }
    });

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
            localStorage.setItem("user_token", token);
            goto("/user/dashboard");
        } catch (err) {
            error = "Something went wrong.";
            console.error(err);
        }
    };
</script>

<Navbar />

<div
    class="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80"
    aria-hidden="true"
>
    <div
        class="relative left-[calc(80%-11rem)] aspect-1155/678 w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-gradient-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]"
        style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
    ></div>
</div>

<div class="min-h-screen flex items-center justify-center bg-transparent p-6">
    <div class="w-full max-w-md bg-white rounded-xl shadow-lg p-10">
        <div class="text-left mb-8">
            <h1 class="text-3xl font-bold text-black">User Login</h1>
            <p class="text-gray-600 text-sm mt-2">Login to vote securely.</p>
        </div>

        <form on:submit|preventDefault={login} class="space-y-6">
            <div>
                <label class="block text-sm font-medium text-gray-700 mb-1"
                    >Citizenship ID</label
                >
                <input
                    type="text"
                    bind:value={citizenship_id}
                    required
                    placeholder="XX-XX-XX-XXXXX"
                    class="input w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black"
                />
            </div>

            <div>
                <label class="block text-sm font-medium text-gray-700 mb-1"
                    >Voter Card ID</label
                >
                <input
                    type="text"
                    bind:value={voter_card_id}
                    required
                    placeholder="XXXXXXXXXX"
                    class="input w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black"
                />
            </div>

            <div>
                <label class="block text-sm font-medium text-gray-700 mb-1"
                    >Password</label
                >
                <input
                    type="password"
                    bind:value={password}
                    required
                    placeholder="Password"
                    class="input w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black"
                />
            </div>

            {#if !setup_done}
                <div class="text-sm text-yellow-600 font-medium">
                    Scan this QR code with your authenticator app:
                </div>
                <img
                    src={qr_url}
                    alt="QR Code"
                    class="mx-auto my-2 rounded-md shadow-md"
                />
            {/if}

            <div>
                <label class="block text-sm font-medium text-gray-700 mb-1"
                    >TOTP Code</label
                >
                <input
                    type="text"
                    bind:value={totp_code}
                    placeholder="TOTP Code"
                    class="input w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black"
                />
                <p class="text-xs text-gray-500 text-right mt-1">
                    Leave empty if logging in for the first time.
                </p>
            </div>

            {#if error}
                <div class="text-sm text-red-500 text-center font-medium">
                    {error}
                </div>
            {/if}

            <button
                type="submit"
                class="btn w-full uppercase tracking-wide font-bold btn-primary hover:bg-gray-800 transition duration-300"
            >
                Login
            </button>
            <a class="text-xs block text-blue-500 text-center mt-1" href="/user/forgotPassword">
                Click Me If You Forgot Your Password 
            </a>

            {#if token}
                <div
                    class="mt-4 p-3 bg-green-100 text-green-700 text-sm rounded"
                >
                    Logged in! Token:<br />
                    <code class="break-all">{token}</code>
                </div>
            {/if}
        </form>
    </div>
</div>
