import { SearchResult } from "../types";

const BASE_URL = "http://localhost:8080";

export const fetchFiles = async (): Promise<string[]> => {
  const res = await fetch(`${BASE_URL}/files`);
  return res.json();
};

export const searchInFile = async (
  fileId: string,
  query: string
): Promise<SearchResult[]> => {
  const res = await fetch(`${BASE_URL}/search`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ file_id: fileId, query }),
  });
  return res.json();
};

export const expandContext = async (
  fileId: string,
  index: number
): Promise<{ context: string }> => {
  const res = await fetch(`${BASE_URL}/expand-context`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ file_id: fileId, index }),
  });
  return res.json();
};
