<script>
    let email = '';
    let password = '';
    let error = '';

    async function submitForm() {
        error = '';
        
        if (!email || !password) {
            error = 'Email and password are required';
            return;
        }

        try {
            const res = await fetch('http://localhost:3000/api/admin/login', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ email, password }),
                credentials: 'same-origin'
            });

            const data = await res.json();

            if (!res.ok) {
                error = data.error || 'Unknown error';
                return;
            }

            // Manually set the cookie
            document.cookie = `admin_token=${data.token}; path=/; Secure; HttpOnly; SameSite=Strict`;

            window.location.href = '/';
        } catch (err) {
            error = 'Network or server error';
        }
    }
</script>

<div class="flex items-center justify-center min-h-screen bg-gray-100">
    <div class="bg-white p-8 rounded-lg shadow-lg w-full max-w-sm">
        <h2 class="text-center text-3xl font-semibold text-gray-800 mb-6">Sign in to your account</h2>
        <form on:submit|preventDefault={submitForm} class="space-y-4">
            <div>
                <label for="email" class="block text-sm font-medium text-gray-600">Email</label>
                <input type="email" bind:value={email} required class="input input-bordered w-full p-3 rounded-md" placeholder="Enter your email" />
            </div>
            <div>
                <label for="password" class="block text-sm font-medium text-gray-600">Password</label>
                <input type="password" bind:value={password} class="input input-bordered w-full p-3 rounded-md" required placeholder="Enter your password" />
            </div>
            <button type="submit" class="btn btn-primary w-full p-3 rounded-md">Sign in</button>
        </form>
        {#if error}
            <div class="text-red-500 text-sm mt-3 text-center">{error}</div>
        {/if}
    </div>
</div>
