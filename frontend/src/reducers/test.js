export default function test(state = 0,action) {
    switch(action.type) {
        case 'TEST_1':
            return state + 1;
        case 'TEST_2':
            return 2;
        default:
            return 0;
    }
}