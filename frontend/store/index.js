import Vuex from "vuex";
import axios from "axios"

const createStore = () => {
    return new Vuex.Store({
        state: {
            fetchedPosts: []
        },
        mutations: {
            setPosts(state, posts) {
                state.fetchedPosts = posts
            },
            addPost(state, post) {
                state.fetchedPosts.push(post)
            },
            updatePost(state, editedPost) {
                console.log("Mutations => " + editedPost);
                let post_index = state.fetchedPosts.findIndex(post => post.id == editedPost.id)
                console.log("Mutations | POST INDEX => " + post_index)
                state.fetchedPosts[post_index] = editedPost
            }
        },
        actions: {
            nuxtServerInit(vuexContext, context) {
                return axios.get("http://localhost:8080/api/task")
                    .then(response => {
                        //console.log(response.data)
                        console.log("nuxt server init")
                        let data = response.data;
                        let postArray = []
                        for (let key in data) {
                            // data["id"] = key
                            postArray.push({id: key, ...data[key]})
                        }
                        vuexContext.commit("setPosts", postArray)
                    })
            },
            setPosts(vuexContext, posts) {

            },
            addPost(vuexContext, post) {

            },
            updatePost(vuexContext, editedPost) {

            }

        },
        getters: {
            getPosts(state) {
                console.log("entered into getPosts ")
                return state.fetchedPosts
            }
        }
    })
}

export default createStore
