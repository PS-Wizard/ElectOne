<script>
    import Footer from "../../../components/footer.svelte";
    import Navbar from "../../../components/Navbar.svelte";

    let citizenship_id = "";
    let voter_card_id = "";
    let password = "";
    let acceptedTerms = false;

    let citizenshipFront = null;
    let citizenshipBack = null;
    let voterCard = null;
    let selfie = null;

    let message = "";

    function isValidCitizenship(id) {
        return /^\d{2}-\d{2}-\d{2}-\d{5}$/.test(id);
    }

    function isValidVoterCard(id) {
        return /^\d{10}$/.test(id);
    }

    function isValidPassword(pw) {
        return (
            pw.length >= 8 &&
            /[A-Z]/.test(pw) &&
            /\d/.test(pw) &&
            /[^A-Za-z0-9]/.test(pw)
        );
    }

    async function submitAppeal() {
        if (!citizenshipFront || !citizenshipBack || !voterCard || !selfie) {
            message = "Please upload all 4 required photos.";
            return;
        }

        if (!isValidCitizenship(citizenship_id)) {
            message = "Invalid Citizenship ID format.";
            return;
        }

        if (!isValidVoterCard(voter_card_id)) {
            message = "Invalid Voter Card ID.";
            return;
        }

        if (!isValidPassword(password)) {
            message =
                "Password must be at least 8 chars, include a number, uppercase, and symbol.";
            return;
        }

        if (!acceptedTerms) {
            message = "Please Accept The Terms And Conditions";
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
                <p
                    class="text-xs text-gray-500 font-bold block text-right mt-1"
                >
                    At least 8 characters and include a number & special
                    character.
                </p>
            </div>

            <div>
                <label class="block text-xs font-bold text-gray-700 mb-1"
                    >Citizenship Front</label
                >
                <input
                    type="file"
                    accept="image/*"
                    required
                    class="file-input border-1 w-full"
                    on:change={(e) => (citizenshipFront = e.target.files[0])}
                />
            </div>

            <div>
                <label class="block text-xs font-bold text-gray-700 mb-1"
                    >Citizenship Back</label
                >
                <input
                    type="file"
                    accept="image/*"
                    required
                    class="file-input border-1 w-full"
                    on:change={(e) => (citizenshipBack = e.target.files[0])}
                />
            </div>

            <div>
                <label class="block text-xs font-bold text-gray-700 mb-1"
                    >Voter Card</label
                >
                <input
                    type="file"
                    accept="image/*"
                    required
                    class="file-input border-1 w-full"
                    on:change={(e) => (voterCard = e.target.files[0])}
                />
            </div>

            <div>
                <label class="block text-xs font-bold text-gray-700 mb-1"
                    >Selfie</label
                >
                <input
                    type="file"
                    accept="image/*"
                    required
                    class="file-input border-1 w-full"
                    on:change={(e) => (selfie = e.target.files[0])}
                />
            </div>

            <!-- After Selfie input -->
            <div class="flex items-center mt-4 justify-start">
                <input
                    type="checkbox"
                    bind:group={acceptedTerms}
                    required
                    class="mr-2 mt-1 checkbox checkbox-sm rounded-sm"
                />
                <label class="text-xs text-gray-600">
                    I agree to the
                    <a
                        href="/docs/tos"
                        target="_blank"
                        class="text-blue-600 underline">Terms & Conditions</a
                    >
                    and
                    <a
                        href="/docs/cookies"
                        target="_blank"
                        class="text-blue-600 underline">Cookie Policy</a
                    >.
                </label>
            </div>

            <div class="flex justify-center">
                <button
                    type="submit"
                    class="btn w-full uppercase tracking-wide font-bold btn-primary hover:bg-gray-800 transition duration-300"
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
<Footer />
