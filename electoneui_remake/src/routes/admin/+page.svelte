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
        const res = await fetch("http://localhost:3000/auth/admin", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ email, password, totp_code }),
        });

        const data = await res.json();

        if (!res.ok) {
            error = data.message || "Login failed";
            return;
        }

        if (!data.setup_done && data.qr_url) {
            setup_done = false;
            qr_url = `https://api.qrserver.com/v1/create-qr-code/?size=200x200&data=${encodeURIComponent(data.qr_url)}`;
        } else {
            token = data.token;
            localStorage.setItem("token", token);
            goto("/admin/dashboard")
        }
    };
</script>

<div class="max-w-sm mx-auto mt-10 p-6 bg-base-200 rounded-xl shadow space-y-4">
    <h1 class="text-xl font-bold">Admin Login</h1>

    <input
        type="email"
        placeholder="Email"
        class="input input-bordered w-full"
        bind:value={email}
    />
    <input
        type="password"
        placeholder="Password"
        class="input input-bordered w-full"
        bind:value={password}
    />

    {#if !setup_done}
        <p class="text-sm text-warning">
            Scan this QR code with an authenticator app:
        </p>
        <img src={qr_url} alt="QR Code" class="mx-auto" />
    {/if}

    <input
        type="text"
        placeholder="TOTP Code"
        class="input input-bordered w-full"
        bind:value={totp_code}
    />

    {#if error}
        <div class="alert alert-error">{error}</div>
    {/if}

    <button class="btn btn-primary w-full" on:click|preventDefault={login}
        >Login</button
    >

    {#if token}
        <div class="alert alert-success">
            Logged in! Token: <br /><code class="break-all">{token}</code>
        </div>
    {/if}
</div>
