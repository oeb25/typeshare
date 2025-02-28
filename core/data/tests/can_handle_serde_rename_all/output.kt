/**
 * Generated by typeshare 1.1.0
 */

@file:NoLiveLiterals

package com.agilebits.onepassword

import androidx.compose.runtime.NoLiveLiterals
import kotlinx.serialization.*

/// This is a Person struct with camelCase rename
@Serializable
data class Person (
	val firstName: String,
	val lastName: String,
	val age: UByte,
	val extraSpecialField1: Int,
	val extraSpecialField2: List<String>? = null
)

/// This is a Person2 struct with UPPERCASE rename
@Serializable
data class Person2 (
	val FIRST_NAME: String,
	val LAST_NAME: String,
	val AGE: UByte
)

