import React, { ReactNode } from "react";
import { cn } from "@/lib/utils";

type Action = {
  label: string;
  onClick: (id: string) => void;
  element?: ReactNode;
}

type ListProps<T> = {
  items: T[];
  renderItem: (item: T, column: string) => ReactNode;
  columnsHeaders: string[];
  excludeKeys?: string[];
  columnKeyMapping: Record<string, string>; // Mapping of headers to keys
  className?: string;
  children?: ReactNode;
  actions?: Action[];

}

const List: React.FC<ListProps<any>> = ({
  children,
  renderItem,
  items,
  className,
  columnsHeaders,
  excludeKeys = [],
  columnKeyMapping,
  actions = [],
  ...otherProps
}) => {
  const listClasses = cn(`min-w-full divide-y divide-gray-200`, className)
  const allKeys = items.length > 0 ? (Object.keys(items[0]) as string[]) : [];
  const filteredKeys = allKeys.filter(key => !excludeKeys.includes(key));

  return (<>
    <table className={listClasses} {...otherProps}>
      <thead className="bg-gray-100 text-gray-600">
        <tr>
          {columnsHeaders.map((column, index) => (
            <th key={index} className="px-6 py-3 text-left font-semibold border-b border-gray-300">{column}</th>
          ))}
        </tr>
      </thead>
      <tbody className="bg-white divide-y divide-gray-200">
        {items.map((item: any, index) => (
          <tr key={index} className="hover:bg-gray-50 transition-colors duration-200">
            {columnsHeaders.map((column, columnIndex) => {
              const mappedKey = columnKeyMapping[column];
              if (mappedKey && filteredKeys.includes(mappedKey)) {
                return (
                  <td key={columnIndex} className="px-6 py-4 whitespace-nowrap text-gray-800 border-b border-gray-300">
                    {renderItem(item, mappedKey)}
                  </td>
                );
              }
              return null;
            })}
            <td>
              {actions.map((action, actionIndex) => {
                return React.cloneElement(
                  action.element as React.ReactElement, {
                  key: actionIndex,
                  onClick: () => action.onClick(item.id),
                }
                );
              })}
            </td>
            {children}
          </tr>
        ))}
      </tbody>
    </table >
  </>
  )
}

export default List;
