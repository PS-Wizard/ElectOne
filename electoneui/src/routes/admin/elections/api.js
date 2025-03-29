const BASE_URL = "https://localhost:3000/api/secure";

export async function GetElections(page = 0) {
    const res = await fetch(`${BASE_URL}/electionsPaginated/${page}`, {
        method: "GET",
        credentials: "include"
    });

    if (!res.ok) {
        console.error('Failed to fetch elections:', res.status);
        return [];
    }

    const data = await res.json();
    console.log("Fetched elections:", data);
    return data;
}

export async function DeleteElection(id) {
    const res = await fetch(`${BASE_URL}/election/${id}`, {
        method: "DELETE",
        credentials: "include"
    });
    return res.json();
}

export async function UpdateElection(electionID, election) {
    const res = await fetch(`${BASE_URL}/election/${electionID}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify(election)
    });
    
    if (!res.ok) {
        throw new Error(`HTTP error! status: ${res.status}`);
    }
    
    // Handle empty response
    const text = await res.text();
    return text ? JSON.parse(text) : {};
}

export async function CreateElection(election) {
    const res = await fetch(`${BASE_URL}/election`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify(election)
    });
    return res.json();
}
