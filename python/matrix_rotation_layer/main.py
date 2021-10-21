def snakeOrder(i: int, rows: int, cols: int) -> tuple[int, int]:
    rows -= 1 # offset for 0 base
    cols -= 1 # offset for 0 base
    i = i % (2*rows + 2*cols) # outer elements of matrix from ring

    if i == 0:
        return 0, 0
    elif i < rows:
        return i, 0
    elif i < rows + cols:
        return rows, i - rows
    elif i < 2*rows + cols:
        return (2*rows + cols) - i, cols
    else:
        return 0, 2*rows + 2*cols - i


def copyMatrix(matrix: list[list[int]]) -> list[list[int]]:
    newMatrix = []
    for i in range(len(matrix)):
        newMatrix.append([])
        for j in range(len(matrix[i])):
            newMatrix[i].append(matrix[i][j])
            
    return newMatrix
        
        
def traverseSnakeOrder(matrix: list[list[int]], rows: int, cols: int, r: int) -> list[list[int]]:
    snakeList = []
    newMatrix = copyMatrix(matrix)
    for i in range(rows*2 + cols*2 - 4):
        snakeList.append(snakeOrder(i, rows, cols))
    
    #print(f"snakeList:{snakeList}")
    for x in range(len(snakeList)):
        i, j = snakeList[x]
        a, b = snakeList[(x - r) % len(snakeList)]
        #print(f"i:{i}, j:{j}, a:{a}, b:{b}, len(snakeList):{len(snakeList)}")
        matrix[i][j] = newMatrix[a][b]
   
    return matrix
        

def peelShell(matrix: list[list[int]]) -> list[list[int]]:
    newMatrix = []
    for i in range(1, len(matrix)-1):
        newMatrix.append(matrix[i][1:-1])
    #print(f"peelShell-{newMatrix}")
    return newMatrix


def updateSubMatrix(matrix: list[list[int]], sub: list[list[int]]) -> None:
    subRows = len(sub)
    subCols = len(sub[0])
    diffRows = (len(matrix) - subRows)//2
    diffCols = (len(matrix[0]) - subCols)//2

    #print(f"Update sub matrix pre update: \nmatrix:")
    #printMatrix(matrix)
    #print(f"sub:")
    #printMatrix(sub)
    
    for i in range(subRows):
        for j in range(subCols):
            matrix[i+diffRows][j+diffCols] = sub[i][j]
    
    #print(f"Update sub matrix post update: \nmatrix:")
    #printMatrix(matrix)
    #print(f"sub:")
    #printMatrix(sub)

    
def printMatrix(matrix: list[list[int]]) -> None:
    for i in range(len(matrix)):
        for j in range(len(matrix[i])):
            print(f"{matrix[i][j]} ", end='')
        print()


def matrixRotation(matrix: list[list[int]], r: int) -> None:
    rows = len(matrix)
    cols = len(matrix[0])
    
    rotMatrix = [[]]
    subMatrix = None
    for i in range(min(rows//2, cols//2)):
        if i == 0:
            rotMatrix = traverseSnakeOrder(matrix, rows, cols, r)
            subMatrix = rotMatrix
        else:
            subMatrix = traverseSnakeOrder(peelShell(matrix), rows-2*i, cols-2*i, r)
            updateSubMatrix(rotMatrix, subMatrix)
    
    printMatrix(matrix)
    

if __name__ == '__main__':
    first_multiple_input = input().rstrip().split()

    m = int(first_multiple_input[0])

    n = int(first_multiple_input[1])

    r = int(first_multiple_input[2])

    matrix = []

    for _ in range(m):
        matrix.append(list(map(int, input().rstrip().split())))

    matrixRotation(matrix, r)
