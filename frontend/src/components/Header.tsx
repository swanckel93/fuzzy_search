import { useStore } from "../store/useStore";
import { useEffect, useState } from "react";
import { fetchFiles, searchInFile } from "../utils/api";

export default function Header() {
  const {
    files,
    currentFile,
    searchQuery,
    setFiles,
    setCurrentFile,
    setSearchQuery,
    setSearchResults,
  } = useStore();

  const [file, setFile] = useState<File | null>(null); // Track the selected file

  useEffect(() => {
    fetchFiles().then(setFiles);
  }, []);

  const handleSearch = async () => {
    if (currentFile && searchQuery) {
      const results = await searchInFile(currentFile, searchQuery);
      setSearchResults(results);
    }
  };

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files) {
      const selectedFile = event.target.files[0];
      setFile(selectedFile);
    }
  };

  const handleFileUpload = async () => {
    if (!file) {
      alert("Please select a file first.");
      return;
    }

    const formData = new FormData();
    formData.append("file", file);

    try {
      const response = await fetch("http://localhost:8080/upload", {
        method: "POST",
        body: formData,
      });

      if (!response.ok) {
        throw new Error("File upload failed");
      }

      alert("File uploaded successfully");
      fetchFiles().then(setFiles); // Reload files after upload
    } catch (error) {
      console.error("Error uploading file:", error);
      alert("Error uploading file");
    }
  };

  return (
    <div className="flex flex-col sm:flex-row items-center gap-2 p-4 bg-gray-100 rounded-b-lg shadow-md">
      {/* Document Selection */}
      <select
        value={currentFile ?? ""}
        onChange={(e) => setCurrentFile(e.target.value)}
        className="p-2 rounded border"
      >
        <option value="" disabled>
          Select a document
        </option>
        {files.map((file) => (
          <option key={file} value={file}>
            {file}
          </option>
        ))}
      </select>

      {/* Search Input */}
      <input
        type="text"
        placeholder="Search..."
        className="p-2 border rounded w-full sm:w-64"
        value={searchQuery}
        onChange={(e) => setSearchQuery(e.target.value)}
        onKeyDown={(e) => e.key === "Enter" && handleSearch()}
      />

      {/* Search Button */}
      <button
        className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
        onClick={handleSearch}
      >
        Search
      </button>

      {/* File Upload Section */}
      <div className="flex items-center gap-4">
        {/* Hidden File Input */}
        <input
          type="file"
          accept=".txt"
          className="hidden"
          id="file-upload"
          onChange={handleFileChange}
        />
        {/* File Upload Button */}
        <label
          htmlFor="file-upload"
          className="cursor-pointer px-4 py-2 bg-blue-500 text-white rounded-lg"
        >
          Select File
        </label>

        {/* Upload Button */}
        <button
          onClick={handleFileUpload}
          className="px-4 py-2 bg-green-500 text-white rounded-lg"
        >
          Upload File
        </button>
      </div>
    </div>
  );
}

