import React, { ReactNode } from "react";

interface ListProps<T> {
  items: T[];
  renderItem: (item: T, column: string) => ReactNode;
  columns: string[]
  className?: string;
}

const List: React.FC<ListProps<any>> = ({ items, renderItem, className, columns, ...otherProps }) => {
  return (
    <table className={`table-auto ${className}`} {...otherProps}>
      <thead>
        <tr>
          {columns.map((column, index) => (
            <th key={index} className="px-4 py-2">{column}</th>
          ))}
        </tr>
      </thead>
      <tbody>
        {items.map((item: any, index) => (
          <tr key={index}>
            {columns.map((column, columnIndex) => (
              <td key={columnIndex} className="px-4 py-2">
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
