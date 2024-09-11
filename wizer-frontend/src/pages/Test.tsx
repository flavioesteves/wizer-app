import User from "@/models/User";
import api from "@/services/api";
import { useEffect, useState } from "react";




const Test: React.FC = () => {
  const [data, setData] = useState<User[] | null>(null);

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const response: User[] | null = await api.user.getAll();

        if (response !== null) {
          setData(response);
        }

      } catch (error) {
        console.error("Error fetching users:", error);
      }
    };
    fetchUsers();
  }, []);



  return (
    <div className="container mx-auto p-4">
      {data && (
        <ul className="list-disc space-y-2">
          {data.map((user: User) => (
            <li key={user.id} className="text-gray-700">
              <h3 className="text-lg font-bold">{user.email}</h3>
              <p>{user.role}</p>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default Test;

