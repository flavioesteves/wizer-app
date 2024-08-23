import React, { useState, useEffect } from "react";
import User from "@/models/User";


interface ListLoaderProps {
  fetchData: () => Promise<User[]>;
}

const ListLoader: React.FC<ListLoaderProps> = ({ fetchData }) => {
  const [data, setData] = useState<User[] | null>(null);

  useEffect(() => {
    const fetchListItems = async () => {
      try {
        const fetchedData = await fetchData();
        setData(fetchedData);
      } catch (error) {
        console.log(error);
      }
    };

    fetchListItems()
  }, [fetchData])

  return (
    <div className="container mx-auto p-4">
      {
        data && (
          <ul className="list-disc space-y-2">
            {data.map((user) => (
              <li key={user.id} className="text-gray-700">
                <h3 className="text-lg font-bold">{user.email}</h3>
                <p>{user.role}</p>
              </li>
            ))}
          </ul>
        )}
    </div >
  );
};

export default ListLoader;
