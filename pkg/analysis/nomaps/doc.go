/*
nomaps provides a linter to ensure that fields do not use map types.

Maps are discouraged in Kubernetes APIs. It is hard to distinguish between structs and maps in JSON/YAML and as such, lists of named subobjects are preferred over plain map types.

Instead of

  ports:
    www:
      containerPort: 80

use

  ports:
    - name: www
      containerPort: 80

Lists should use the `+listType=map` and `+listMapKey=name` markers, or equivalent.
*/

package nomaps
