import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faPlus } from "@fortawesome/free-solid-svg-icons";

import api from "@/services/api";
import { Routine, ColumnKeyMapping } from "@/models/Routine";
import List from "@/components/ui/list";
import Button from "@/components/ui/button";

const renderRoutineItem = (routine: Routine, key: string) => {
  switch (key) {
    case "name":
      return (<p className="text-blue-500 font-semibold">{routine.name}</p>);
    case "profile_id":
      return (<p className="text-blue-500 font-semibold">{routine.profile_id}</p>);
    case "exercises":
      return (
        <div className="flex space-x-2 text-blue-500 font-semibold">
          {routine.exercises.map((exercise, index) => (
            <div key={index} className="flex-1">
              <p>{exercise.name}</p>
            </div>
          ))}
        </div>
      );
    default:
      return "";
  }
}

const RoutineList = () => {
  const navigate = useNavigate();
  const col = Object.keys(ColumnKeyMapping);
  const excludedKeys: (keyof Routine)[] = ["id"];
  const [routines, setRoutines] = useState<Routine[]>([]);

  useEffect(() => {
    const fetchRoutines = async () => {
      try {
        const response: Routine[] | null = await api.routine.getAll();
        if (response) {
          setRoutines(response);
        }
      } catch (error) {
        console.error("Error fetching routines:", error);
      }
    };
    fetchRoutines();
  }, []);

  const handleEdit = (id: string) => {
    navigate(`/routines/${id}`)
  }

  const handleDelete = (id: string) => {
    if (window.confirm("Are you sure, you want to delete this routine?")) {
      api.routine.deleteRoutine(id).then(() => {
        setRoutines((prev) => prev.filter((routine) => routine.id !== id));
      }).catch((error) =>
        console.error("Error deleting routine:", error));
    }
  }

  const actions = [
    {
      label: "Edit",
      onClick: handleEdit,
      element: <Button variant="primary">Edit</Button>
    },
    {
      label: "Delete",
      onClick: handleDelete,
      element: <Button variant="red">Delete</Button>
    }

  ]


  return (
    <>
      <h3 className="text-2xl font-bold mb-4 text-center">Routines</h3>
      <div className="flex flex-row justify-center overflow-x-auto">
        <div className="flex justify-center px-4 flex-grow">
          <List
            excludeKeys={excludedKeys}
            items={routines}
            renderItem={renderRoutineItem}
            actions={actions}
            columnsHeaders={col}
            columnKeyMapping={ColumnKeyMapping}
            className="border rounded-lg p-2">
          </List>
        </div>
        <Link to="/routines/new">
          <Button variant="green" className="rounded-full">
            <FontAwesomeIcon icon={faPlus} size="lg" className="fa-fw rounded-full" />
            New
          </Button>
        </Link>
      </div >
    </>
  )

}

export default RoutineList;
