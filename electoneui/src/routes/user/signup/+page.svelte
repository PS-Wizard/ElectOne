<script>
    let citizenship_id = "";
    let voter_card_id = "";
    let password = "";

    let citizenshipFront = null;
    let citizenshipBack = null;
    let voterCard = null;
    let selfie = null;

    let message = "";

    async function submitAppeal() {
        if (!citizenshipFront || !citizenshipBack || !voterCard || !selfie) {
            message = "Please upload all 4 required photos.";
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

<form on:submit|preventDefault={submitAppeal}>
    <input
        type="text"
        placeholder="Citizenship ID"
        bind:value={citizenship_id}
        required
    />
    <input
        type="text"
        placeholder="Voter Card ID"
        bind:value={voter_card_id}
        required
    />
    <input
        type="password"
        placeholder="Password"
        bind:value={password}
        required
    />

    <label>Citizenship ID (Front)</label>
    <input
        type="file"
        accept="image/*"
        on:change={(e) => (citizenshipFront = e.target.files[0])}
        required
    />

    <label>Citizenship ID (Back)</label>
    <input
        type="file"
        accept="image/*"
        on:change={(e) => (citizenshipBack = e.target.files[0])}
        required
    />

    <label>Voter Card</label>
    <input
        type="file"
        accept="image/*"
        on:change={(e) => (voterCard = e.target.files[0])}
        required
    />

    <label>Selfie</label>
    <input
        type="file"
        accept="image/*"
        on:change={(e) => (selfie = e.target.files[0])}
        required
    />

    <button type="submit">Submit Appeal</button>
</form>

{#if message}
    <p>{message}</p>
{/if}
