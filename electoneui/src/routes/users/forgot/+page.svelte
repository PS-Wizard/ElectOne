<script>
    let otp = "";
</script>


<div class="relative flex min-h-screen flex-col justify-center overflow-hidden bg-gray-50 py-12 text-black">
  <div class="relative bg-white px-6 pt-10 pb-9 shadow-xl mx-auto w-full max-w-lg rounded-2xl">
    <div class="mx-auto flex w-full max-w-md flex-col space-y-8">
      <div class="flex flex-col items-center justify-center text-center space-y-2">
        <div class="font-semibold text-3xl">
          <p>Authenticator App Verification</p>
        </div>
        <div class="flex flex-row text-sm font-medium text-black">
          <p>We’ve sent a code to your authenticator app. Enter it below.</p>
        </div>
      </div>

      <div>
        <form action="" method="post">
          <div class="flex justify-center space-x-4">
            <div class="relative">
              <input
                type="text"
                maxlength="1"
                bind:value={otp[0]}
                class="w-16 h-16 text-3xl text-center border-2 border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
                inputmode="numeric"
              />
              <div class="absolute top-0 left-0 flex items-center justify-center w-6 h-6 rounded-full bg-blue-600 text-white text-xs font-semibold">
                1
              </div>
            </div>

            <div class="relative">
              <input
                type="text"
                maxlength="1"
                bind:value={otp[1]}
                class="w-16 h-16 text-3xl text-center border-2 border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
                inputmode="numeric"
              />
              <div class="absolute top-0 left-0 flex items-center justify-center w-6 h-6 rounded-full bg-blue-600 text-white text-xs font-semibold">
                2
              </div>
            </div>

            <div class="relative">
              <input
                type="text"
                maxlength="1"
                bind:value={otp[2]}
                class="w-16 h-16 text-3xl text-center border-2 border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
                inputmode="numeric"
              />
              <div class="absolute top-0 left-0 flex items-center justify-center w-6 h-6 rounded-full bg-blue-600 text-white text-xs font-semibold">
                3
              </div>
            </div>

            <div class="relative">
              <input
                type="text"
                maxlength="1"
                bind:value={otp[3]}
                class="w-16 h-16 text-3xl text-center border-2 border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
                inputmode="numeric"
              />
              <div class="absolute top-0 left-0 flex items-center justify-center w-6 h-6 rounded-full bg-blue-600 text-white text-xs font-semibold">
                4
              </div>
            </div>

            <div class="relative">
              <input
                type="text"
                maxlength="1"
                bind:value={otp[4]}
                class="w-16 h-16 text-3xl text-center border-2 border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
                inputmode="numeric"
              />
              <div class="absolute top-0 left-0 flex items-center justify-center w-6 h-6 rounded-full bg-blue-600 text-white text-xs font-semibold">
                5
              </div>
            </div>

            <div class="relative">
              <input
                type="text"
                maxlength="1"
                bind:value={otp[5]}
                class="w-16 h-16 text-3xl text-center border-2 border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
                inputmode="numeric"
              />
              <div class="absolute top-0 left-0 flex items-center justify-center w-6 h-6 rounded-full bg-blue-600 text-white text-xs font-semibold">
                6
              </div>
            </div>
          </div>

          <div class="mt-8">
            <button
              type="submit"
              class="w-full py-3 text-lg font-semibold text-white bg-blue-600 rounded-lg shadow-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              Verify
            </button>
          </div>

          <div class="mt-4 text-center text-sm text-gray-500">
            <p>Didn't receive a code?</p>
            <a href="#" class="text-blue-600 hover:underline">Resend</a>
          </div>
        </form>
      </div>
    </div>
  </div>
</div>

