<script>
    import Navbar from "../../../components/Navbar.svelte";
    // import Footer from "../../../components/footer.svelte";

    let citizenship_id = "";
    let voter_card_id = "";
    let password = "";

    let citizenshipFront = null;
    let citizenshipBack = null;
    let voterCard = null;
    let selfie = null;

    let message = "";

    const passwordRegex = /^(?=.*[0-9])(?=.*[!@#$%^&*])(?=.{8,})/;

    async function submitAppeal() {
        if (!citizenshipFront || !citizenshipBack || !voterCard || !selfie) {
            message = "Please upload all 4 required photos.";
            return;
        }

        if (!passwordRegex.test(password)) {
            message =
                "Password must be at least 8 characters and include a number and a special character.";
            return;
        }

        const formData = new FormData();
        formData.append("citizenship_id", citizenship_id);
        formData.append("voter_card_id", voter_card_id);
        formData.append("password", password);
        formData.append("photos", citizenshipFront);
        formData.append("photos", citizenshipBack);
        formData.append("photos", voterCard);
        formData.append("photos", selfie);

        try {
            const res = await fetch("http://localhost:3000/appeal/", {
                method: "POST",
                body: formData,
            });
            const data = await res.json();

            if (!res.ok) {
                message = data.message || "Something went wrong.";
                return;
            }

            message = `Appeal submitted! ID: ${data.appeal_id}`;
        } catch (err) {
            message = "Error: " + err.message;
        }
    }
</script>

<Navbar />

<div class="min-h-screen flex items-center justify-center bg-gray-100 p-6">
    <div class="w-full max-w-md bg-white rounded-xl shadow-lg p-10">
        <div class="text-center mb-8">
            <h1 class="text-3xl font-bold text-black">Create Appeal</h1>
            <p class="text-gray-600 text-sm mt-2">
                Submit your appeal details.
            </p>
        </div>

        <form
            on:submit|preventDefault={submitAppeal}
            enctype="multipart/form-data"
            class="space-y-6"
        >
            <div>
                <label class="block text-sm font-medium text-gray-700 mb-1"
                    >Citizenship ID</label
                >
                <input
                    type="text"
                    bind:value={citizenship_id}
                    required
                    placeholder="123-456"
                    class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black"
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
                    placeholder="VC-001"
                    class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black"
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
                    placeholder="securepassword123"
                    class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black"
                />
                <p class="text-xs text-gray-500 mt-1">
                    Must be at least 8 characters and include a number & special
                    character.
                </p>
            </div>

            <div>
                <label class="block text-sm font-medium text-gray-700 mb-1"
                    >Citizenship Front</label
                >
                <input
                    type="file"
                    accept="image/*"
                    required
                    on:change={(e) => (citizenshipFront = e.target.files[0])}
                    class="w-full"
                />
            </div>

            <div>
                <label class="block text-sm font-medium text-gray-700 mb-1"
                    >Citizenship Back</label
                >
                <input
                    type="file"
                    accept="image/*"
                    required
                    on:change={(e) => (citizenshipBack = e.target.files[0])}
                    class="w-full"
                />
            </div>

            <div>
                <label class="block text-sm font-medium text-gray-700 mb-1"
                    >Voter Card</label
                >
                <input
                    type="file"
                    accept="image/*"
                    required
                    on:change={(e) => (voterCard = e.target.files[0])}
                    class="w-full"
                />
            </div>

            <div>
                <label class="block text-sm font-medium text-gray-700 mb-1"
                    >Selfie</label
                >
                <input
                    type="file"
                    accept="image/*"
                    required
                    on:change={(e) => (selfie = e.target.files[0])}
                    class="w-full"
                />
            </div>

            <div class="flex justify-center">
                <button
                    type="submit"
                    class="bg-black text-white text-lg font-medium py-3 px-10 rounded-full hover:bg-gray-800 transition duration-300"
                >
                    Submit Appeal
                </button>
            </div>

            {#if message}
                <p class="text-sm text-center mt-4 text-red-500">{message}</p>
            {/if}
        </form>
    </div>
</div>

<!-- <Footer /> -->
