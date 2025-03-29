<script>
    let email = "";
    let password = "";
    let errorMessage = "";

    const handleLogin = async () => {
        try {
            const response = await fetch(
                "https://localhost:3000/api/admin/login",
                {
                    method: "POST",
                    credentials: "include",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        email: email,
                        password: password,
                    }),
                },
            );

            const data = await response.json();

            if (response.ok) {
                // Handle successful login
                console.log(data.message); // Redirect or show success
            } else {
                // Display error
                errorMessage = data.error || "Unknown error occurred";
            }
        } catch (error) {
            errorMessage = "Network error, please try again.";
        }
    };
</script>

<div class="min-h-screen flex justify-center items-center">
    <div class="p-8 rounded-lg shadow-lg w-96">
        <h1 class="text-2xl font-semibold text-center mb-6">Admin Login</h1>

        <!-- Error Message -->
        {#if errorMessage}
            <div class="text-red-600 text-sm mb-4">
                {errorMessage}
            </div>
        {/if}

        <div>
            <label for="email" class="block text-gray-700">Email</label>
            <input
                id="email"
                type="email"
                class="input w-full p-3 border border-gray-300 rounded-md mt-2"
                bind:value={email}
                placeholder="Enter your email"
            />

            <label for="password" class="block text-gray-700 mt-4"
                >Password</label
            >
            <input
                id="password"
                type="password"
                class="input w-full p-3 border border-gray-300 rounded-md mt-2"
                bind:value={password}
                placeholder="Enter your password"
            />

            <button
                on:click={handleLogin}
                class="cursor-pointer w-full bg-white text-black p-3 rounded-md mt-6 border border-transparent hover:bg-black hover:text-white hover:border-white transition"
            >
                Login
            </button>
        </div>
    </div>
</div>
