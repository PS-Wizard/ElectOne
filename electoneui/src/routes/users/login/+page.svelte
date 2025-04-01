<script>
    let citizenID = "";
    let password = "";
    let error = "";

    const login = async () => {
        error = "";
        if (!citizenID || !password) {
            error = "Citizen ID and password are required";
            return;
        }

        try {
            const res = await fetch("https://localhost:3000/api/userlogin", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ citizenID, password }),
            });

            if (!res.ok) throw new Error("Invalid credentials");
        } catch (err) {
            error = err.message;
        }
    };
</script>

<div
    class="flex flex-col items-center justify-center min-h-screen bg-gray-100 p-4"
>
    <div class="card w-96 bg-white shadow-lg p-6 rounded-lg">
        <h2 class="text-xl font-semibold text-center mb-4 text-black">Login</h2>
        <input
            type="text"
            placeholder={error || "Citizen ID"}
            bind:value={citizenID}
            class="input bg-white text-black input-bordered w-full mb-2 placeholder-red-500"
        />
        <input
            type="password"
            placeholder={error && !citizenID ? "Password" : "Password"}
            bind:value={password}
            class="input bg-white text-black input-bordered w-full mb-4 placeholder-red-500"
        />
        <button on:click={login} class="btn btn-primary w-full">Login</button>
        <a
            href="/forgot-password"
            class="text-sm text-blue-500 text-center block mt-2"
            >Forgot Password?</a
        >
    </div>
</div>
