# GPS injector
This cli tool can:
    - download the ls_node collection and add coordinated fields. The output is saved in the specified file. ``gps-injector pull -o <file.json>``
    - upload the coordinates into a new collection called ls_node_coordinates ``gps-injector push -i <file.json>``
    - to overwrite an existing ls_node_coordinates collection use ``gps-injector push --i <file.json> -o=true``
## adjust jalapeno parameter
update the environment file to match your Jalapeno ArangoDB connection parameters