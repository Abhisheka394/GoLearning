# The Tricolor Algorithm

This Tricolor Marks and Sweep algorithm divides objects in the heap into three sets according to the color assigned, such as black, white, and gray. The white set objects are the candidates for garbage collection.

White set  - objects to be cleaned
Gray set - objects under observation
Black set  - objects to keep

The black set objects guarantee to have no reference to any objects in the white set. But note: objects in the white set may have reference to the objects in the black set; this has no effect on the operation of the garbage collector. The objects in the gray set might have reference to the objects in the white set.

The steps in the Tricolor Scheme can grossly be summed up as follows:

Step 1: Pick an object from the gray set.
Step 2: Gray all the object references and move it to the black set
Step 3: Repeat step 1 and 2 until the gray set is empty.
Step 4: Put all other objects in the black set if they are reachable from the root otherwise put them into the white set.
Step 5: The objects in the white set can be garbage collected.
