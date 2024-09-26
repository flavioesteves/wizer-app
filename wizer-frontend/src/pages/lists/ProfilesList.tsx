import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faPlus } from "@fortawesome/free-solid-svg-icons";

import api from "@/services/api";
import List from "@/components/ui/list";
import Button from "@/components/ui/button";
import { ColumnKeyMapping, Profile } from "@/models/Profile";



const renderProfileItem = (profile: Profile, key: string) => {
  switch (key) {
    case "user_id":
      return (<p className="text-blue-500 font-semibold">{profile.user_id}</p>)
    case "gender":
      return (<p className="text-blue-500 font-semibold">{profile.gender}</p>)
    case "birth_year":
      return (<p className="text-blue-500 font-semibold">{profile.birth_year}</p>)
    case "height_cm":
      return (<p className="text-blue-500 font-semibold">{profile.height_cm}</p>)
    case "weight_kg":
      return (<p className="text-blue-500 font-semibold">{profile.weight_kg}</p>)
    case "body_fat_percentage":
      return (<p className="text-blue-500 font-semibold">{profile.body_fat_percentage}</p>)
    case "goal":
      return (<p className="text-blue-500 font-semibold">{profile.goal}</p>)
    default:
      return "";
  };
}

const ProfileList = () => {
  const navigate = useNavigate();
  const excludedKeys: (keyof Profile)[] = ["id"];
  const col = Object.keys(ColumnKeyMapping);
  const [profiles, setProfiles] = useState<Profile[]>([]);


  useEffect(() => {
    const fetchProfiles = async () => {
      try {
        const response: Profile[] | null = await api.profile.getAll();
        if (response) {
          setProfiles(response)
        }
      } catch (error) {
        console.error("Error fetching profiles:", error)
      }
    };
    fetchProfiles();
  }, []);


  const handleEdit = (id: string) => {
    navigate(`/profiles/${id}`)
  }

  const handleDelete = (id: string) => {
    console.log(id)
    //TODO: Create modal components
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
  ];


  return (
    <>
      <h3 className="text-2xl font-bold mb-4 text-center">Profiles</h3>
      <div className="flex flex-row justify-center">
        <div className="flex justify-center px-4 flex-grow">
          <List
            excludeKeys={excludedKeys}
            columnKeyMapping={ColumnKeyMapping}
            items={profiles}
            renderItem={renderProfileItem}
            actions={actions}
            columnsHeaders={col}
            className="border rounded-lg p-2">
          </List>
        </div>
        <Link to="/profiles/new">
          <Button variant="green" className="rounded-full">
            <FontAwesomeIcon icon={faPlus} size="lg" className="fa-fw rounded-full" />
            New
          </Button>
        </Link>
      </div>
    </>
  )
}


export default ProfileList;



