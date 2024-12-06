#include <array>
#include <fstream>
#include <iostream>
#include <vector>
using namespace std;

enum MOVE_RESULT{OUT, OK, STUCK, LOOP};
struct Direction {
    int x;
    int y;
    bool operator==(const Direction &d) const {
        return x == d.x && y == d.y;
    }
    [[nodiscard]] bool isOpposite(const Direction &d) const {
        return x + d.x + y + d.y == 0;
    }
};

struct Position {
    int x;
    int y;

    [[nodiscard]] Position move(const Direction d) const {

        return {x + d.x, y+ d.y};
    }

};
struct Result {
    int nbPos;
    int isLoop;
};
array<Direction, 4> directions = {Direction{-1,0}, {0, 1}, {1,0}, {0,-1}};

Result moveInMap(vector<vector<char>> map, Position pos, Direction d) ;

int main() {
    vector<vector<char>> array;
    ifstream inputFile("../input");
    Position startPos = {0, 0};

    if (inputFile.is_open()) {
        string line;
        while (getline(inputFile, line)) {
            vector<char> row;
            for (char c : line) {
                row.push_back(c);
            }
            array.push_back(row);
        }
        inputFile.close();

        for (int i = 0; i < array.size(); i++) {
            for (int j = 0; j < array[i].size(); j++) {
                if (array[i][j] == '^') {
                    startPos = {i, j};
                    array[i][j] = 'X';

                }
            }
        }

    } else {
        cerr << "Error opening file." << endl;
    }
    // part 1
    cout << "Part 1 : " << moveInMap(array, startPos, directions[0]).nbPos << endl;

    int nbLoop = 0;
    for (int i = 0; i < array.size(); i++) {
        for (int j = 0; j < array[i].size(); j++) {
            if (array[i][j] == '.') {
                vector<vector<char>> newArray = array;
                newArray[i][j] = '#';
                if (auto [nbPos, isLoop] = moveInMap(newArray, startPos, directions[0]); isLoop) {
                    nbLoop++;
                }
            }
        }
    }
    cout << "Part 2 : " << nbLoop << endl;

    return 0;
}
MOVE_RESULT checkMove(const vector<vector<char>> &map, const Position &pos, const Direction &d, int nbIter) {
    auto [x, y] = pos.move(d);
    if (x < 0 || y < 0 || x >= map.size() || y >= map[0].size()) {
        return OUT;
    }
    if (map.at(x).at(y) == '#') {
        return STUCK;
    }
    if (nbIter > map.size() * map[0].size()) {
        return LOOP;
    }
    return OK;
}



Result moveInMap(vector<vector<char>> map, Position pos, Direction d) {

    int nbPos = 1;
    size_t dirIndex = 0;
    bool end = false;
    int nbIter = 0;
    while (!end) {
        switch (checkMove(map, pos, d, nbIter)) {
            case OUT:
                end = true;
                break;
            case OK:
                pos = pos.move(d);
                nbIter++;

                if (map[pos.x][pos.y] == '.') {
                    map[pos.x][pos.y] = 'X';
                    nbPos++;
                }

                break;
            case STUCK:
                dirIndex = (dirIndex + 1)%4;
                d = directions[dirIndex];
                break;
            case LOOP:
                return {nbPos, true};
            default:
                break;
        }
    };


    return {nbPos, false};
}


