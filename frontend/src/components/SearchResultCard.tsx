import { useState } from "react";
import { expandContext } from "../utils/api";
import { useStore } from "../store/useStore";
import { ArrowDown, ArrowUp } from "lucide-react";

interface Props {
  sentence: string;
  index: number;
}

export default function SearchResultCard({ sentence, index }: Props) {
  const { currentFile, searchQuery } = useStore();

  const [sentences, setSentences] = useState<string[]>([sentence]);
  const [startIndex, setStartIndex] = useState(index);
  const [endIndex, setEndIndex] = useState(index);

  const handleExpandUp = async () => {
    const newIndex = startIndex - 1;
    if (!currentFile || newIndex < 0) return;

    const { context } = await expandContext(currentFile, newIndex);
    setSentences((prev) => [context, ...prev]);
    setStartIndex(newIndex);
  };

  const handleExpandDown = async () => {
    const newIndex = endIndex + 1;
    if (!currentFile) return;

    const { context } = await expandContext(currentFile, newIndex);
    setSentences((prev) => [...prev, context]);
    setEndIndex(newIndex);
  };

  const handleReset = () => {
    setSentences([sentence]);
    setStartIndex(index);
    setEndIndex(index);
  };

  const highlightMatch = (text: string, query: string) => {
    if (!query) return text;
    const regex = new RegExp(`(${query})`, 'gi');
    return text.split(regex).map((part, i) =>
      part.toLowerCase() === query.toLowerCase() ? (
        <span key={i} className="bg-yellow-300">{part}</span>
      ) : (
        part
      )
    );
  };

  return (
    <div className="bg-white p-4 rounded-2xl shadow mb-4 relative">
      <button
        onClick={handleExpandUp}
        className="absolute top-2 right-4 text-gray-500 hover:text-gray-800"
        title="Expand previous"
      >
        <ArrowUp size={18} />
      </button>

      <div className="prose max-w-none space-y-2 mt-6">
        {sentences.map((s, i) => (
          <p key={i}>{highlightMatch(s, searchQuery)}</p>
        ))}
      </div>

      <div className="mt-4 flex justify-between items-center">
        <button
          onClick={handleReset}
          className="text-sm px-3 py-1 bg-red-100 text-red-700 rounded hover:bg-red-200"
        >
          Reset
        </button>

        <button
          onClick={handleExpandDown}
          className="text-gray-500 hover:text-gray-800"
          title="Expand next"
        >
          <ArrowDown size={18} />
        </button>
      </div>
    </div>
  );
}
