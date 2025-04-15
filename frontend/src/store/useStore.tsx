import { create } from "zustand";
import { SearchResult } from "../types";

interface StoreState {
  files: string[];
  currentFile: string | null;
  searchQuery: string;
  searchResults: SearchResult[];
  setFiles: (files: string[]) => void;
  setCurrentFile: (file: string) => void;
  setSearchQuery: (query: string) => void;
  setSearchResults: (results: SearchResult[]) => void;
}

export const useStore = create<StoreState>((set) => ({
  files: [],
  currentFile: null,
  searchQuery: "",
  searchResults: [],
  setFiles: (files) => set({ files }),
  setCurrentFile: (file) => set({ currentFile: file }),
  setSearchQuery: (query) => set({ searchQuery: query }),
  setSearchResults: (results) => set({ searchResults: results }),
}));
