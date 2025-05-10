<script>
    import Footer from "../components/footer.svelte";
    import Navbar from "../components/Navbar.svelte";
</script>

<section>
    <Navbar />
    <section class="w-screen overflow-hidden">
        <div
            class="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80"
            aria-hidden="true"
        >
            <div
                class="relative left-[calc(50%-11rem)] aspect-1155/678 w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-linear-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]"
                style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
            ></div>
        </div>
        <div class="bg-transparent text-black">
            <div class="relative isolate px-6 pt-14 lg:px-8">
                <div class="mx-auto max-w-2xl py-32 sm:py-48 lg:py-56">
                    <div class="hidden sm:mb-8 sm:flex sm:justify-center">
                        <div
                            class="relative rounded-full px-3 py-1 text-sm/6 text-gray-600 ring-1 ring-gray-900/10 hover:ring-gray-900/20"
                        >
                            Announcing the new voting system for a better
                            democracy .
                            <a href="/" class="font-semibold text-black"
                                ><span
                                    class="absolute inset-0"
                                    aria-hidden="true"
                                ></span>Read more
                                <span aria-hidden="true">&rarr;</span></a
                            >
                        </div>
                    </div>
                    <div class="text-center">
                        <h1
                            class="text-5xl font-semibold tracking-tight text-black sm:text-7xl"
                        >
                            A New Era of Voting for Nepal's Future
                        </h1>
                        <p
                            class="mt-8 text-lg font-medium text-gray-400 sm:text-xl/8"
                        >
                            Empowering citizens with a transparent, secure, and
                            accessible voting system for a stronger democracy.
                        </p>
                        <div
                            class="mt-10 flex items-center justify-center gap-x-6"
                        >
                            <a
                                href="/user/signup"
                                class="text-sm/6 font-semibold bg-black text-white px-5 py-2 rounded-lg"
                                >Get Started <span aria-hidden="true"
                                    >&rarr;</span
                                ></a
                            >
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
        </div>

        <!-- Features Section -->
        <div class="bg-transparent py-24 sm:py-32" id="features">
            <div class="mx-auto max-w-7xl px-6 lg:px-8">
                <div class="mx-auto max-w-2xl lg:text-center">
                    <h2 class="text-base/7 font-semibold text-indigo-600">
                        Streamlined Voting
                    </h2>
                    <p
                        class="mt-2 text-4xl font-semibold tracking-tight text-gray-900 sm:text-5xl lg:text-balance"
                    >
                        Everything you need for secure and transparent elections
                    </p>
                    <p class="mt-6 text-lg/8 text-gray-600">
                        Empowering voters with an easy-to-use platform for
                        democratic participation.
                    </p>
                </div>
                <div
                    class="mx-auto mt-16 max-w-2xl sm:mt-20 lg:mt-24 lg:max-w-4xl"
                >
                    <dl
                        class="grid max-w-xl grid-cols-1 gap-x-8 gap-y-10 lg:max-w-none lg:grid-cols-2 lg:gap-y-16"
                    >
                        <div class="relative pl-16">
                            <dt class="text-base/7 font-semibold text-gray-900">
                                <div
                                    class="absolute top-0 left-0 flex size-10 items-center justify-center rounded-lg bg-indigo-600"
                                >
                                    <svg
                                        class="size-6 text-white"
                                        fill="none"
                                        viewBox="0 0 24 24"
                                        stroke-width="1.5"
                                        stroke="currentColor"
                                        aria-hidden="true"
                                    >
                                        <path
                                            stroke-linecap="round"
                                            stroke-linejoin="round"
                                            d="M12 16.5V9.75m0 0 3 3m-3-3-3 3M6.75 19.5a4.5 4.5 0 0 1-1.41-8.775 5.25 5.25 0 0 1 10.233-2.33 3 3 0 0 1 3.758 3.848A3.752 3.752 0 0 1 18 19.5H6.75Z"
                                        />
                                    </svg>
                                </div>
                                Secure Voting
                            </dt>
                            <dd class="mt-2 text-base/7 text-gray-600">
                                Ensuring transparency and integrity in every
                                vote cast. No tampering or fraud allowed.
                            </dd>
                        </div>
                        <div class="relative pl-16">
                            <dt class="text-base/7 font-semibold text-gray-900">
                                <div
                                    class="absolute top-0 left-0 flex size-10 items-center justify-center rounded-lg bg-indigo-600"
                                >
                                    <svg
                                        class="size-6 text-white"
                                        fill="none"
                                        viewBox="0 0 24 24"
                                        stroke-width="1.5"
                                        stroke="currentColor"
                                        aria-hidden="true"
                                    >
                                        <path
                                            stroke-linecap="round"
                                            stroke-linejoin="round"
                                            d="M16.5 10.5V6.75a4.5 4.5 0 1 0-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 0 0 2.25-2.25v-6.75a2.25 2.25 0 0 0-2.25-2.25H6.75a2.25 2.25 0 0 0-2.25 2.25v6.75a2.25 2.25 0 0 0 2.25 2.25Z"
                                        />
                                    </svg>
                                </div>
                                Real-time Results
                            </dt>
                            <dd class="mt-2 text-base/7 text-gray-600">
                                Live tracking of votes, providing instant
                                results once voting ends. Keep track of your
                                election progress as it happens.
                            </dd>
                        </div>
                        <div class="relative pl-16">
                            <dt class="text-base/7 font-semibold text-gray-900">
                                <div
                                    class="absolute top-0 left-0 flex size-10 items-center justify-center rounded-lg bg-indigo-600"
                                >
                                    <svg
                                        class="size-6 text-white"
                                        fill="none"
                                        viewBox="0 0 24 24"
                                        stroke-width="1.5"
                                        stroke="currentColor"
                                        aria-hidden="true"
                                    >
                                        <path
                                            stroke-linecap="round"
                                            stroke-linejoin="round"
                                            d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99"
                                        />
                                    </svg>
                                </div>
                                Voter Privacy
                            </dt>
                            <dd class="mt-2 text-base/7 text-gray-600">
                                We prioritize voter confidentiality, ensuring
                                that every vote remains anonymous and secure.
                            </dd>
                        </div>
                        <div class="relative pl-16">
                            <dt class="text-base/7 font-semibold text-gray-900">
                                <div
                                    class="absolute top-0 left-0 flex size-10 items-center justify-center rounded-lg bg-indigo-600"
                                >
                                    <svg
                                        class="size-6 text-white"
                                        fill="none"
                                        viewBox="0 0 24 24"
                                        stroke-width="1.5"
                                        stroke="currentColor"
                                        aria-hidden="true"
                                    >
                                        <path
                                            stroke-linecap="round"
                                            stroke-linejoin="round"
                                            d="M7.864 4.243A7.5 7.5 0 0 1 19.5 10.5c0 2.92-.556 5.709-1.568 8.268M5.742 6.364A7.465 7.465 0 0 0 4.5 10.5a7.464 7.464 0 0 1-1.15 3.993m1.989 3.559A11.209 11.209 0 0 0 8.25 10.5a3.75 3.75 0 1 1 7.5 0c0 .527-.021 1.049-.064 1.565M12 10.5a14.94 14.94 0 0 1-3.6 9.75m6.633-4.596a18.666 18.666 0 0 1-2.485 5.33"
                                        />
                                    </svg>
                                </div>
                                Election Integrity
                            </dt>
                            <dd class="mt-2 text-base/7 text-gray-600">
                                Automated systems with constant monitoring to
                                ensure no interference with the election
                                process.
                            </dd>
                        </div>
                    </dl>
                </div>
            </div>
            <div class="blur-3xl">
                <div
                    class="relative left-[calc(20%+3rem)] aspect-1155/678 w-[36.125rem] -translate-x-1/2 bg-linear-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%+36rem)] sm:w-[72.1875rem]"
                    style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
                ></div>
            </div>
        </div>

        <!-- FAQ Section -->
        <div class="bg-transparent py-24 sm:py-32 w-full" id="about">
            <div class="mx-auto max-w-7xl px-6 lg:px-8">
                <div class="mx-auto max-w-2xl lg:text-center">
                    <h2 class="text-base/7 font-semibold text-indigo-600">
                        Frequently Asked Questions
                    </h2>
                    <p
                        class="mt-2 text-4xl font-semibold tracking-tight text-gray-900 sm:text-5xl lg:text-balance"
                    >
                        Everything you need to know about the election system
                    </p>
                </div>

                <div class="mt-16 max-w-1xl sm:mt-20 lg:mt-24 text-black">
                    <div
                        class="collapse collapse-plus border-2 border-black mb-4"
                    >
                        <input
                            type="radio"
                            name="my-accordion"
                            checked="checked"
                        />
                        <div class="collapse-title font-semibold">
                            How do I vote?
                        </div>
                        <div class="collapse-content text-sm">
                            Voting is simple! After registering, log in to your
                            account and select the election you want to
                            participate in. Follow the on-screen prompts to cast
                            your vote.
                        </div>
                    </div>

                    <div
                        class="collapse collapse-plus border-2 border-black mb-4"
                    >
                        <input type="radio" name="my-accordion" />
                        <div class="collapse-title font-semibold">
                            Is my vote anonymous?
                        </div>
                        <div class="collapse-content text-sm">
                            Yes! We use advanced encryption and privacy measures
                            to ensure that all votes are anonymous and cannot be
                            traced back to you.
                        </div>
                    </div>

                    <div
                        class="collapse collapse-plus text-black border-2 border-black mb-4"
                    >
                        <input type="radio" name="my-accordion" />
                        <div class="collapse-title font-semibold">
                            Can I change my vote?
                        </div>
                        <div class="collapse-content text-sm">
                            Once your vote is submitted, it cannot be changed.
                            Be sure to review your selection carefully before
                            confirming your vote.
                        </div>
                    </div>

                    <div
                        class="collapse collapse-plus text-black border-2 border-black mb-4"
                    >
                        <input type="radio" name="my-accordion" />
                        <div class="collapse-title font-semibold">
                            What if I encounter a problem during voting?
                        </div>
                        <div class="collapse-content text-sm">
                            If you face any issues while voting, please contact
                            our support team immediately via the contact page.
                            We'll be happy to assist you.
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <Footer />
    </section>
</section>
