<script>
    let citizenID = "";
    let password = "";
    let otp = ""; // OTP input value
    let error = "";
    let totpSecret = ""; // This will hold the TOTP secret URL for QR code

    const login = async () => {
        error = "";
        if (!citizenID || !password) {
            error = "Citizen ID and password are required";
            return;
        }

        try {
            const res = await fetch("https://localhost:3000/api/userlogin", {
                method: "POST",
                credentials: "include",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ citizenID, password }),
            });

            if (!res.ok) throw new Error("Invalid credentials");

            // If TOTP secret is returned, update the state
            const data = await res.json();
            if (data.totp_secret) {
                totpSecret = data.totp_secret; // Store the TOTP secret URL
                const modal = document.getElementById("my_modal_3");
                modal.showModal(); // Show the modal with QR code
            } else {
                // No need for QR code, just prompt for OTP
                showOTPField(); // Display OTP input field if TOTP is already set
            }
        } catch (err) {
            error = err.message;
            console.log(error);
        }
    };

    const showOTPField = () => {
        const modal = document.getElementById("my_modal_3");
        modal.showModal(); // Show the OTP input modal
    };

    const submitOTP = async () => {
        if (!otp) {
            error = "OTP is required";
            return;
        }

        try {
            const res = await fetch(
                `https://localhost:3000/api/secure/verifyOTP?otp=${otp}&citizenID=${citizenID}`,
                {
                    method: "GET",
                    credentials: "include",
                },
            );

            if (!res.ok) throw new Error("Invalid OTP");

            // Navigate to the dashboard or another page if OTP is valid
            window.location.href = "/users"; // Example: redirect to dashboard
        } catch (err) {
            error = err.message;
            console.log(error);
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
            href="/users/forgot"
            class="text-sm text-blue-500 text-center block mt-2"
            >Forgot Password?</a
        >
    </div>
</div>

<!-- Modal for QR Code or OTP -->
<dialog id="my_modal_3" class="modal">
    <div class="modal-box">
        <form method="dialog">
            <button
                class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2"
                >✕</button
            >
        </form>
        <h3 class="text-lg font-bold">
            Please scan this code or enter the OTP
        </h3>
        {#if totpSecret}
            <!-- Show QR code image for new users -->
            <img
                src={`https://localhost:3000/api/secure/qr/${totpSecret}`}
                alt="Scan this QR code"
                class="mb-4"
            />
            <p class="py-4">
                Please scan this QR code with your authenticator app to complete
                the login.
            </p>
        {/if}
        <input
            type="text"
            placeholder="Enter OTP"
            bind:value={otp}
            class="input bg-white text-black input-bordered w-full mb-2"
        />
        <button on:click={submitOTP} class="btn btn-primary w-full"
            >Submit OTP</button
        >
    </div>
</dialog>
