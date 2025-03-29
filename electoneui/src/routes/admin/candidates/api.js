const BASE_URL = "https://localhost:3000/api/secure";

export async function GetCandidates(page = 0) {
    const res = await fetch(`${BASE_URL}/candidatesPaginated/${page}`, {
        method: "GET",
        credentials: "include"
    });
    return res.ok ? res.json() : [];
}

export async function CreateCandidate(candidate) {
    const res = await fetch(`${BASE_URL}/candidate`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify(candidate)
    });
    return res.json();
}

export async function UpdateCandidate(id, candidate) {
    const res = await fetch(`${BASE_URL}/candidate/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify(candidate)
    });
    return res.json();
}

export async function DeleteCandidate(id) {
    const res = await fetch(`${BASE_URL}/candidate/${id}`, {
        method: "DELETE",
        credentials: "include"
    });
    return res.json();
}
