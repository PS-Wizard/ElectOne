<script>
    import Footer from "../../../components/footer.svelte";
    import Navbar from "../../../components/Navbar.svelte";
    // Function to go back to the previous page
</script>

<Navbar />
<!-- Full Page Layout -->
<div class="min-h-screen flex flex-col">
    <div
        class="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80"
        aria-hidden="true"
    >
        <div
            class="relative left-[calc(50%-11rem)] aspect-1155/678 w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-linear-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]"
            style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
        ></div>
    </div>
    <!-- Top Navbar with Back Arrow and Title -->
    <div class="flex items-center justify-between px-6 py-4">
        <!-- Back Arrow with no navigation -->
        <button
            on:click={() => window.history.back()}
            class="text-xl text-black font-bold cursor-pointer">←</button
        >
    </div>

    <!-- Terms & Conditions Content -->
    <div class="w-full text-gray-700 py-10 px-6">
        <h1 class="text-6xl font-bold text-gray-900 text-center mb-6">
            Terms & Conditions
        </h1>
        <p class="text-sm text-gray-600 mb-6 text-center">
            By using this Online Voting System, you agree to the terms below.
            Please read carefully.
        </p>

        <!-- Terms Sections -->
        <div class="space-y-6 text-base text-left max-w-5xl mx-auto">
            <div>
                <h2 class="text-xl font-semibold text-gray-800 mb-1">
                    1. Eligibility
                </h2>
                <p>
                    Only verified and eligible Nepali citizens may vote through
                    this system using valid credentials such as Voter ID and
                    National Identity.
                </p>
            </div>

            <div>
                <h2 class="text-xl font-semibold text-gray-800 mb-1">
                    2. Use of System
                </h2>
                <p>
                    The platform must be used solely for legitimate voting.
                    Misuse, including multiple voting attempts or tampering, is
                    strictly prohibited and may result in legal action.
                </p>
            </div>

            <div>
                <h2 class="text-xl font-semibold text-gray-800 mb-1">
                    3. Data Privacy
                </h2>
                <p>
                    All personal information, including voter details, is <span
                        class="text-red-600 font-semibold"
                        >securely encrypted</span
                    > and will not be shared except as required by law.
                </p>
            </div>

            <div>
                <h2 class="text-xl font-semibold text-gray-800 mb-1">
                    4. Vote Confidentiality
                </h2>
                <p>
                    Every vote cast is <span class="text-red-600 font-semibold"
                        >anonymous</span
                    > and cannot be traced back to the voter, ensuring free and fair
                    elections.
                </p>
            </div>

            <div>
                <h2 class="text-xl font-semibold text-gray-800 mb-1">
                    5. System Availability
                </h2>
                <p>
                    We strive for 24/7 availability during elections but are not
                    responsible for technical issues due to unforeseen events or
                    network problems.
                </p>
            </div>

            <div>
                <h2 class="text-xl font-semibold text-gray-800 mb-1">
                    6. Administrator Rights
                </h2>
                <p>
                    Administrators reserve the right to start, end, or cancel
                    elections and to take action against users who violate these
                    terms.
                </p>
            </div>

            <div>
                <h2 class="text-xl font-semibold text-gray-800 mb-1">
                    7. Modifications
                </h2>
                <p>
                    Terms may be updated at any time. Users will be notified of
                    changes through official channels or app notifications.
                </p>
            </div>

            <div>
                <h2 class="text-xl font-semibold text-gray-800 mb-1">
                    8. Legal Compliance
                </h2>
                <p>
                    All users must comply with the Election Commission of
                    Nepal’s regulations and relevant laws under the Nepalese
                    Constitution and Cyber Law.
                </p>
            </div>
        </div>
    </div>
    <div
        class="absolute inset-x-0 top-[calc(100%-13rem)] -z-10 transform-gpu overflow-hidden blur-3xl sm:top-[calc(100%-30rem)]"
        aria-hidden="true"
    >
        <div
            class="relative left-[calc(50%+3rem)] aspect-1155/678 w-[36.125rem] -translate-x-1/2 bg-linear-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%+36rem)] sm:w-[72.1875rem]"
            style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
        ></div>
    </div>
</div>

<Footer />
