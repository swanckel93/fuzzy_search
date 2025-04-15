import { useState } from "react";
import { expandContext } from "../utils/api";
import { useStore } from "../store/useStore";

interface Props {
  sentence: string;
  index: number;
}

export default function SearchResultCard({ sentence, index }: Props) {
  const { currentFile, searchQuery } = useStore();
  const [content, setContent] = useState(sentence);
  const [original] = useState(sentence);

  const handleExpand = async () => {
    if (currentFile) {
      const { context } = await expandContext(currentFile, index);
      setContent(context);
    }
  };

  const handleReset = () => {
    setContent(original);
  };

  // Function to highlight the search query inside the sentence
  const highlightMatch = (sentence: string, query: string) => {
    if (!query) return sentence; // If no query, just return the sentence as is
    const regex = new RegExp(`(${query})`, 'gi'); // Create case-insensitive regex
    return sentence.split(regex).map((part, index) =>
      part.toLowerCase() === query.toLowerCase() ? (
        <span key={index} className="bg-yellow-300">{part}</span> // Highlight matching part
      ) : (
        part // Leave other parts unchanged
      )
    );
  };

  return (
    <div className="bg-white p-4 rounded-2xl shadow mb-4">
      <p className="prose max-w-none">
        {highlightMatch(content, searchQuery)} {/* Highlighting occurs here */}
      </p>
      <div className="mt-2 flex gap-2">
        <button
          onClick={handleExpand}
          className="text-sm px-3 py-1 bg-gray-200 rounded hover:bg-gray-300"
        >
          Expand Context
        </button>
        <button
          onClick={handleReset}
          className="text-sm px-3 py-1 bg-red-100 text-red-700 rounded hover:bg-red-200"
        >
          Reset Context
        </button>
      </div>
    </div>
  );
}
