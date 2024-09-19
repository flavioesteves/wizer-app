import React, { ReactNode } from "react";
import { cn } from "@/lib/utils";

type ListProps<T> = {
  items: T[];
  renderItem: (item: T, column: string) => ReactNode;
  columns: string[]
  className?: string;
}



const List: React.FC<ListProps<any>> = ({ items, renderItem, className, columns, ...otherProps }) => {
  const listClasses = cn(`min-w-full divide-y divide-gray-200`, className)

  console.log("LIST columns: " + columns);
  console.log("LIST renderItem: " + renderItem)
  console.log("LIST items: " + items)

  return (
    <table className={listClasses} {...otherProps}>
      <thead className="bg-gray-100 text-gray-600">
        <tr>
          {columns.map((column, index) => (
            <th key={index} className="px-6 py-3 text-left font-semibold border-b border-gray-300">{column}</th>
          ))}
        </tr>
      </thead>
      <tbody className="bg-white divide-y divide-gray-200">
        {items.map((item: any, index) => (
          <tr key={index} className="hover:bg-gray-50 transition-colors duration-200">
            {columns.map((column, columnIndex) => (
              <td key={columnIndex} className="px-6 py-4 whitespace-nowrap text-gray-800 border-b border-gray-300">
                {renderItem(item, column)}
              </td>
            ))}
          </tr>
        ))}
      </tbody>
    </table>
  )
}

export default List;
