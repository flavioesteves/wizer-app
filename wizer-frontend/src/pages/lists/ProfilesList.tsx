import List from "@/components/ui/list";
import { ColumnKeyMapping } from "@/models/Profile";

const renderProfileItem = (profile: { id: number; name: string; user: string }, column: string) => {
  switch (column) {
    case "Id":
      return (<p className="text-gray-700 font-semibold">{profile.id}</p>)
    case "Name":
      return (<p className="text-blue-500 font-semibold">{profile.name}</p>)
    case "User":
      return (<p className="text-blue-500 font-semibold">{profile.user}</p>)
    default:
      return "";
  };
}
// <div className="flex justify-between items-center">
//   <p className="text-gray-700">{profile.id}</p>
//   <p className="text-blue-500">{profile.name}</p>
// </div>

const ProfileList = () => {
  const columns = ["Id", "Name", "User", "Created", "Updated"]
  const emptyItems = Array.from({ length: 2 }).map(() => ({}));
  const profiles = [
    { id: 1, name: "test 1", user: "user 1" },
    { id: 2, name: "test 2", user: "user 2" },
  ]

  const diplayedProfiles = [...profiles, ...emptyItems].slice(0, 20);



  return (
    <>
      <h3 className="text-2xl font-bold mb-4 text-center">Profiles</h3>
      <div className="flex flex-row justify-center">
        <div className="flex justify-center px-4 flex-grow">
          <List
            columnKeyMapping={ColumnKeyMapping}
            items={diplayedProfiles}
            renderItem={renderProfileItem}
            columnsHeaders={columns}
            className="border rounded-lg p-2"
          />
        </div>
      </div>
    </>
  )
}


export default ProfileList;



