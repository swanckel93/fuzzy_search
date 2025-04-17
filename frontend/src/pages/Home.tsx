import { useStore } from "../store/useStore";
import SearchResultCard from "../components/SearchResultCard";
import Header from "../components/Header";

export default function Home() {
  const { searchResults } = useStore();

  return (
    <div className="min-h-screen bg-gray-50 p-4">
      <Header />
      <div className="mt-6 space-y-4">
        {searchResults
          .slice() // clone to avoid mutating store state directly
          .sort((a, b) => {
            if (a.distance !== b.distance) return a.distance - b.distance;
            return a.index - b.index;
          })
          .map((res) => (
            <SearchResultCard
              key={`${res.index}-${res.distance}-${res.match}`}
              sentence={res.sentence}
              index={res.index}
            />
          ))}
      </div>
    </div>
  );
}
